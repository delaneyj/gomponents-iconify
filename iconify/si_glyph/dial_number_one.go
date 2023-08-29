package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialNumberOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><circle cx="2" cy="3" r="2"/><circle cx="7" cy="3" r="2"/><circle cx="13" cy="2" r="2"/><circle cx="2" cy="8" r="2"/><circle cx="7" cy="8" r="2"/><circle cx="12" cy="8" r="2"/><circle cx="2" cy="13" r="2"/><circle cx="7" cy="13" r="2"/><circle cx="12" cy="13" r="2"/></g>`),
		g.Group(children),
	)
}