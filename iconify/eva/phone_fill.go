package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPhoneFill0"><g id="evaPhoneFill1"><path id="evaPhoneFill2" fill="currentColor" d="M17.4 22A15.42 15.42 0 0 1 2 6.6A4.6 4.6 0 0 1 6.6 2a3.94 3.94 0 0 1 .77.07a3.79 3.79 0 0 1 .72.18a1 1 0 0 1 .65.75l1.37 6a1 1 0 0 1-.26.92c-.13.14-.14.15-1.37.79a9.91 9.91 0 0 0 4.87 4.89c.65-1.24.66-1.25.8-1.38a1 1 0 0 1 .92-.26l6 1.37a1 1 0 0 1 .72.65a4.34 4.34 0 0 1 .19.73a4.77 4.77 0 0 1 .06.76A4.6 4.6 0 0 1 17.4 22Z"/></g></g>`),
		g.Group(children),
	)
}