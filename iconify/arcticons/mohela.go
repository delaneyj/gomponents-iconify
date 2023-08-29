package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mohela(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M31.574 20.965v6.07h3.034M8 27.028v-6.063l3.035 6.07l3.034-6.06v6.06m7.33-6.07v6.07m4.021-6.07v6.07m-4.021-3.046h4.021m1.678.011h1.978m1.056 3.035h-3.034v-6.07h3.034"/><rect width="4.021" height="6.069" x="15.724" y="20.965" rx="2.01" ry="2.01"/><path d="M39.334 25.024h-2.689m-.666 2.011l2.011-6.07l2.01 6.07"/></g><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}