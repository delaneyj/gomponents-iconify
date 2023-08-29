package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetMicOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 23v-2h7v-1h-4v-8h4v-1q0-2.9-2.05-4.95T12 4Q9.1 4 7.05 6.05T5 11v1h4v8H3v-9q0-1.85.713-3.488T5.65 4.65q1.225-1.225 2.863-1.938T12 2q1.85 0 3.488.713T18.35 4.65q1.225 1.225 1.938 2.863T21 11v12h-9Zm-7-5h2v-4H5v4Zm12 0h2v-4h-2v4ZM5 18h2h-2Zm12 0h2h-2Z"/>`),
		g.Group(children),
	)
}