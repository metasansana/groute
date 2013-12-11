package groute_test

import (
	"github.com/metasansana/groute"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Route", func() {

	var (
		flag        int
		route       groute.Route
		server      *httptest.Server
		staticPath  string = "/index123"
		regexPath   string = "/index[0-9]+"
		invalidPath string = "/~/users2"
	)

	flagFunc := func(w http.ResponseWriter, r *http.Request) {

		flag++

	}

	httpFunc := func(aroute string) {

		_, err := http.Get(aroute)

		if err != nil {

			log.Fatal(err)
		}

	}
	DoTests := func() {

		It("should ignore its callback when its route is unmatched", func() {

			httpFunc(server.URL + invalidPath)
			Expect(flag).Should(Equal(0), "Unmatched routes should not execute their callbacks!")
		})

		It("should trigger its callback when its route is matched", func() {

			httpFunc(server.URL + staticPath)
			Expect(flag).Should(Equal(1), "Matched routes should execute their callbacks!")

		})

	}

	BeforeEach(func() {

		flag = 0

		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			route.Trigger(r.URL.Path, w, r)

		}))

	})

	Describe(" when using the route subsystem", func() {

		Context("when StaticRoute is used", func() {
			BeforeEach(func() {

				route = groute.NewStaticRoute(staticPath, flagFunc)

				DoTests()

			})

			Context("when RegexRoute is used", func() {
				BeforeEach(func() {

					route = groute.NewRegexRoute(regexPath, flagFunc)

				})

				DoTests()

			})
		})

	})

})
