package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emoticon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.981.084a7.939 7.939 0 1 0 0 15.876a7.938 7.938 0 0 0 0-15.876zM10.998 4a2 2 0 1 1 .003 3.999A2 2 0 0 1 10.998 4zM5 4a2 2 0 1 1-.001 4.001A2 2 0 0 1 5 4zm3 10c-3.013 0-5-1.899-5-4h10c0 2.101-1.988 4-5 4z"/>`),
		g.Group(children),
	)
}