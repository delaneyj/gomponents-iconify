package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyOldboyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.8 38.511a28.044 28.044 0 0 1-18.594 4.865a1.707 1.707 0 0 1-1.396-1.073l-8.283-31.58a1.075 1.075 0 0 1 .763-1.311l18.605-4.877a1.075 1.075 0 0 1 1.31.763l8.272 31.57a1.6 1.6 0 0 1-.677 1.644Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.177 22.672l12.215-3.2l-2.903-10.921l-12.173 3.248l2.861 10.873zm16.847-2.697c.15.664-.35 1.065-.927 1.529v-.001a31.07 31.07 0 0 1-18.201 4.773c-.702-.075-1.353-.226-1.53-.89m7.505 9.624l-1.118-4.276m-1.578 2.696l4.275-1.117"/><circle cx="29.657" cy="30.659" r="1.096" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.149" cy="28.446" r="1.096" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.887 37.92l1.59-.419m1.697-.44l1.579-.419"/>`),
		g.Group(children),
	)
}