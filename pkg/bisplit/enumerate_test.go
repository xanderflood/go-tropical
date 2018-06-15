package bisplit_test

import (
  "github.com/xanderflood/go-tropical/pkg/bisplit"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Enumerate", func() {
  var (
    W uint
    B uint

    enum bisplit.Enumerator
  )

  JustBeforeEach(func() {
    enum = bisplit.Enumerate(W, B)
  })

  var assertAlwaysNormal = func() {
    It("only produces normalized splits", func() {
      calc := bisplit.ModuliSpace{W: W, B: B}

      for bs := range enum {
        Expect(calc.IsNormal(bs)).To(BeTrue())
      }
    })
  }

  var assertNoSplits = func() {
    It("should not produce any bisplits", func() {
      for range enum {
        Expect(false).To(BeTrue())
      }

      Expect(true).To(BeTrue())
    })
  }

  Context("when there aren't any splits", func() {
    Context("when W=0", func() {
      BeforeEach(func() { W = 0 })

      Context("when B=0", func() {
        BeforeEach(func() { B = 0 })
        assertNoSplits()
      })

      Context("when B=2", func() {
        BeforeEach(func() { B = 2 })
        assertNoSplits()
      })

      Context("when B=3", func() {
        BeforeEach(func() { B = 3 })
        assertNoSplits()
      })

      Context("when B=100", func() {
        BeforeEach(func() { B = 100 })
        assertNoSplits()
      })
    })

    Context("when W=1", func() {
      BeforeEach(func() { W = 1 })

      Context("when B=0", func() {
        BeforeEach(func() { B = 0 })
        assertNoSplits()
      })

      Context("when B=2", func() {
        BeforeEach(func() { B = 2 })
        assertNoSplits()
      })

      Context("when B=3", func() {
        BeforeEach(func() { B = 3 })
        assertNoSplits()
      })

      Context("when B=100", func() {
        BeforeEach(func() { B = 100 })
        assertNoSplits()
      })
    })
  })

  Context("when W=2", func() {
    BeforeEach(func() { W = 2 })

    var assertWAlwaysOne = func() {
      It("only produces bisplits with W = 1", func() {
        for bs := range enum {
          Expect(bs.W).To(Equal(uint(1)))
        }
      })
    }

    Context("when B=2", func() {
      BeforeEach(func() { B = 2 })

      assertAlwaysNormal()
      assertWAlwaysOne()

      It("produces one split", func() {
        splits := []bisplit.Bisplit{}

        for bs := range enum {
          splits = append(splits, bs)
        }

        Expect(len(splits)).To(Equal(2))
        Expect(splits[0]).To(Equal(bisplit.Bisplit{1, 1}))
      })
    })

    Context("when B=3", func() {
      BeforeEach(func() { B = 3 })

      assertAlwaysNormal()
      assertWAlwaysOne()

      It("produces six splits", func() {
        splits := []bisplit.Bisplit{}

        for bs := range enum {
          splits = append(splits, bs)
        }

        Expect(len(splits)).To(Equal(6))
        Expect(splits[0]).To(Equal(
          bisplit.New(
            []bool{true, false},
            []bool{true, false, false},
          ),
        ))
        Expect(splits[1]).To(Equal(
          bisplit.New(
            []bool{true, false},
            []bool{false, true, false},
          ),
        ))
        Expect(splits[2]).To(Equal(
          bisplit.New(
            []bool{true, false},
            []bool{true, true, false},
          ),
        ))
        Expect(splits[5]).To(Equal(
          bisplit.New(
            []bool{true, false},
            []bool{false, true, true},
          ),
        ))
      })
    })

    Context("when B=5", func() {
      BeforeEach(func() { B = 5 })

      assertAlwaysNormal()
      assertWAlwaysOne()

      It("produces 30 splits", func() {
        splits := []bisplit.Bisplit{}

        for bs := range enum {
          splits = append(splits, bs)
        }

        Expect(len(splits)).To(Equal(30))
      })
    })

    Context("when B=10", func() {
      BeforeEach(func() { B = 10 })

      assertAlwaysNormal()
      assertWAlwaysOne()

      It("produces 1022 splits", func() {
        splits := []bisplit.Bisplit{}

        for bs := range enum {
          splits = append(splits, bs)
        }

        Expect(len(splits)).To(Equal(1022))
      })
    })
  })

  Context("when W=3", func() {
    BeforeEach(func() { W = 3 })

    Context("when B=0", func() {
      BeforeEach(func() { B = 0 })
      assertNoSplits()
    })

    Context("when B=1", func() {
      BeforeEach(func() { B = 1 })
      assertNoSplits()
    })

    Context("when B=2", func() {
      BeforeEach(func() { B = 2 })

      It("only produces bisplits with W = 1", func() {
        for bs := range enum {
          Expect(bs.W).To(BeNumerically("<=", uint(3)))
        }
      })
    })
  })

  Context("when W=5", func() {
    BeforeEach(func() { W = 5 })

    Context("when B=5", func() {
      BeforeEach(func() { B = 5 })

      //there are (2^w-2)*(2^b-2)/2
      It("produces 450 non-equivalent splits", func() {
        splits := []bisplit.Bisplit{}
        calc := bisplit.ModuliSpace{W: W, B: B}

        for bs := range enum {
          for _, bs2 := range splits {
            Expect(calc.Equivalent(bs, bs2)).To(BeFalse())
          }

          splits = append(splits, bs)
        }

        Expect(len(splits)).To(Equal(450))
      })
    })
  })
})
