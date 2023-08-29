package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDropUp0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDropUp1" fill="currentColor"><path id="feDropUp2" d="m12 8l7 8H5z"/></g></g>`),
		g.Group(children),
	)
}