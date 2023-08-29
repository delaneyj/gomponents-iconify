package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileScreenShareSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 15v-4.5h4v-2l3 3l-3 3v-2h-2V15H9Zm-4 8V1h14v22H5Zm2-5h10V6H7v12Z"/>`),
		g.Group(children),
	)
}