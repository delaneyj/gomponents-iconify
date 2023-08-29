package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUpSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21V8l7-7l1.85 1.85L15.55 8H23v4.4L19.35 21H8ZM6 8v13H2V8h4Z"/>`),
		g.Group(children),
	)
}