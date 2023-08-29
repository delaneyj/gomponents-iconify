package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filmstrip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h1v2h2V4h9v2h2V4h1v17h-1v-2h-2v2H7v-2H5v2H4V4Zm3 3H5v3h2V7Zm0 4H5v3h2v-3Zm0 4H5v3h2v-3Zm9 3h2v-3h-2v3Zm0-4h2v-3h-2v3Zm0-4h2V7h-2v3ZM8 5v7h7V5H8Zm0 8v7h7v-7H8Z"/>`),
		g.Group(children),
	)
}