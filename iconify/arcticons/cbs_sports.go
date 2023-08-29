package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CbsSports(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="19.706" r="15.206"/><path d="M24 11.325c-6.402 0-12.01 3.35-15.206 8.38c3.196 5.032 8.804 8.38 15.206 8.38s12.01-3.348 15.206-8.38c-3.196-5.03-8.804-8.38-15.206-8.38Z"/><circle cx="24" cy="19.706" r="8.38"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M14.352 43.45V37.4h1.98c1.12 0 2.027.91 2.027 2.032s-.907 2.031-2.027 2.031h-1.98m11.101 1.987V37.4h1.98c1.12 0 2.028.91 2.028 2.032s-.908 2.031-2.027 2.031h-1.98m1.98.001l1.98 1.984m.83-6.047h4.007m-2.004 6.049v-6.049M8.914 42.787c.371.483.836.663 1.483.663h.896a1.51 1.51 0 0 0 1.509-1.51v-.006c0-.833-.676-1.508-1.51-1.508h-.987a1.51 1.51 0 0 1-1.51-1.511h0c0-.836.677-1.514 1.513-1.514h.89c.648 0 1.113.18 1.484.663m22.636 4.723c.37.483.836.663 1.483.663h.896c.833 0 1.508-.676 1.508-1.51v-.006c0-.833-.675-1.508-1.508-1.508h-.988a1.51 1.51 0 0 1-1.51-1.511h0c0-.836.677-1.514 1.513-1.514h.89c.648 0 1.113.18 1.484.663"/><rect width="4.007" height="6.049" x="19.824" y="37.451" rx="2.004" ry="2.004"/></g>`),
		g.Group(children),
	)
}