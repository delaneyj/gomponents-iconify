package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(4)"><ellipse cx="4.469" cy="11.488" rx="1.469" ry="1.488"/><path d="M8.301.021H.779C.371.021.04.326.04.7v14.595c0 .375.33.679.739.679h7.522c.41 0 .739-.304.739-.679V.7c0-.374-.33-.679-.739-.679zM4.5 14a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5zM8 6H1V1h7v5z"/></g>`),
		g.Group(children),
	)
}