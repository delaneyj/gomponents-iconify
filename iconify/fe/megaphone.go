package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Megaphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMegaphone0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMegaphone1" fill="currentColor" fill-rule="nonzero"><path id="feMegaphone2" d="m4 8l.04 2.117L19.5 6.065A1.997 1.997 0 0 1 22 8v8a1.995 1.995 0 0 1-2.5 1.934L4 13.882V16a1 1 0 0 1-2 0V8a1 1 0 1 1 2 0Zm2 3.664v.68L20 16V8L6 11.664Z"/></g></g>`),
		g.Group(children),
	)
}