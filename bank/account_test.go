package bank

import (
    "testing"
    "github.com/onsi/ginkgo"
    "github.com/onsi/gomega"
)

func TestAccount(t *testing.T) {
    gomega.RegisterFailHandler(ginkgo.Fail)
    ginkgo.RunSpecs(t, "Bank Account Management System")
}

var _ = ginkgo.Describe("Bank Account", func() {
    var account Account

    ginkgo.BeforeEach(func() {
        account = Account{}
    })

    ginkgo.Describe("Deposit", func() {
        ginkgo.It("should deposit money into the account", func() {
            err := account.Deposit(100)
            gomega.Expect(err).To(gomega.BeNil())
            gomega.Expect(account.GetBalance()).To(gomega.Equal(100.0))
        })

        ginkgo.It("should return an error for negative deposit", func() {
            err := account.Deposit(-50)
            gomega.Expect(err).To(gomega.HaveOccurred())
        })

        ginkgo.It("should return an error for zero deposit", func() {
            err := account.Deposit(0)
            gomega.Expect(err).To(gomega.HaveOccurred())
        })
    })

    ginkgo.Describe("Withdraw", func() {
        ginkgo.It("should withdraw money from the account", func() {
            account.Deposit(100)
            err := account.Withdraw(50)
            gomega.Expect(err).To(gomega.BeNil())
            gomega.Expect(account.GetBalance()).To(gomega.Equal(50.0))
        })

        ginkgo.It("should return an error for insufficient funds", func() {
            err := account.Withdraw(100)
            gomega.Expect(err).To(gomega.HaveOccurred())
        })

        ginkgo.It("should return an error for negative withdrawal", func() {
            err := account.Withdraw(-50)
            gomega.Expect(err).To(gomega.HaveOccurred())
        })

        ginkgo.It("should return an error for zero withdrawal", func() {
            err := account.Withdraw(0)
            gomega.Expect(err).To(gomega.HaveOccurred())
        })
    })
})
