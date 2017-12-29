package actions_test

import (
	. "cognizant.com/codeblue/actions"
	"os"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"encoding/base64"
	"path/filepath"
)

var _ = Describe("Agent", func() {
	var agent Agent

	BeforeEach(func() {
		agent = NewAgent()
	})

	Describe("ReadFile", func(){
		Context("Non-Exist File", func(){
			It("should throw an error", func(){
				_, err := agent.ReadFile("/tmp/do-not-exist")
				Expect(err).To(Not(BeNil()))
			})
		})

		Context("A valid file", func(){
			It("should return the context of the file", func(){
				dir, err := os.Getwd()
				Expect(err).To(BeNil())

				path := filepath.Join(dir, "cmd.go")
				output, err := agent.ReadFile(path)

				Expect(err).To(BeNil())
				Expect(len(output)).To(BeNumerically(">", 1))
			})
		})
	})

	Describe("DeleteFile", func(){
		Context("Non-Exist File", func(){
			It("should throw an error", func(){
				err := agent.DeleteFile("/tmp/do-not-exist")
				Expect(err).To(Not(BeNil()))
			})
		})
	})

	Describe("StoreFile", func(){
		Context("Empty File", func(){
			It("creates an empty file", func(){
				agent.StoreFile("/tmp/test_empty", "")
				_, err := os.Stat("/tmp/test_empty")
				Expect(err).To(BeNil())
				err = agent.DeleteFile("/tmp/test_empty")
				Expect(err).To(BeNil())
			})
		})

		Context("Base64 File Content", func(){
			It("creates an empty file", func(){
				content := "this is a test"
				contentBase64 := base64.StdEncoding.EncodeToString([]byte(content))
				agent.StoreFile("/tmp/test_content", contentBase64)
				_, err := os.Stat("/tmp/test_content")
				Expect(err).To(BeNil())

				file, err := os.Open("/tmp/test_content")
				Expect(err).To(BeNil())

				stat, err := file.Stat()
				Expect(err).To(BeNil())

				Expect(int(stat.Size())).To(Equal(14))

				err = agent.DeleteFile("/tmp/test_content")
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("RunCmd", func(){
		Context("Empty Cmd", func(){
			It("should ignore empty spaces", func(){
				output, err := agent.RunCmd(" ")
				Expect(err).To(BeNil())
				Expect(len(output)).To(Equal(0))
			})

			It("should do nothing", func(){
				output, err := agent.RunCmd("")
				Expect(err).To(BeNil())
				Expect(len(output)).To(Equal(0))
			})
		})

		Context("Without argument", func(){
			It("should run the command without error", func(){
				_, err := agent.RunCmd("ls")
				Expect(err).To(BeNil())
			})

			It("should ignore empty spaces", func(){
				output, err := agent.RunCmd("   ls  ")
				Expect(err).To(BeNil())
				Expect(output).To(Not(BeNil()))
			})

			It("should capture output", func(){
				output, err := agent.RunCmd("ls")
				Expect(err).To(BeNil())
				Expect(output).To(Not(BeNil()))
			})
		})

		Context("With arguments", func(){
			It("should run the command without error", func(){
				output, err := agent.RunCmd("ls -a -l")
				Expect(err).To(BeNil())
				Expect(output).To(Not(BeNil()))
			})

			It("provide accurate output", func(){
				output, err := agent.RunCmd("echo something else")
				Expect(err).To(BeNil())
				Expect(output).To(Not(BeNil()))
				Expect(string(output)).To(Equal("something else\n"))
			})

			Context("With quote", func(){
				It("provide accurate output", func(){
					output, err := agent.RunCmd("echo 'something else'")
					Expect(err).To(BeNil())
					Expect(output).To(Not(BeNil()))
					Expect(string(output)).To(Equal("something else\n"))
				})
			})
		})
	})
})
