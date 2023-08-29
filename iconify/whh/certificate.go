package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Certificate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1024 389l-91 122l91 121q-4 18-7 29l-140 60l17 151q-9 10-20 21l-151-18l-61 140l-28 7l-122-91l-122 91q-18-4-28-7l-61-140l-151 18q-7-8-20-21l17-151L7 661q-3-10-7-29l91-121L0 389q4-17 7-28l140-60l-17-151q9-10 20-21l151 18L362 7q9-3 28-7l122 90L634 0q18 4 28 7l61 140l151-18q10 10 20 21l-17 151l140 60z"/>`),
		g.Group(children),
	)
}