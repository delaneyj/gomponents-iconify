package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUpOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21V8l7-7l1.85 1.85L14.55 8H23v4.4L19.35 21H7Zm2-2h9l3-7v-2h-9l1.35-5.5L9 8.85V19ZM9 8.85V19V8.85ZM7 8v2H4v9h3v2H2V8h5Z"/>`),
		g.Group(children),
	)
}