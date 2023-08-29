package simple_line_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M974.944 65.392c0-23.232-12.592-44.654-32.912-55.935a63.765 63.765 0 0 0-31.088-8.063a63.96 63.96 0 0 0-33.775 9.648L141.44 457.634c-15.952 9.905-26.512 26.208-29.376 44.4V32.004c0-17.664-14.336-32-32-32s-32 14.336-32 32v960c0 17.664 14.336 32 32 32s32-14.336 32-32V521.939c2.88 18.208 13.44 34.511 29.375 44.384l736.72 446.64a63.881 63.881 0 0 0 33.776 9.664a63.937 63.937 0 0 0 31.088-8.065a63.958 63.958 0 0 0 32.912-55.936zM175.2 511.985L910.944 65.393l1.008 893.216z"/>`),
		g.Group(children),
	)
}