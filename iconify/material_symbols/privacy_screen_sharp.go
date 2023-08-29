package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivacyScreenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11.6L9.6 4H2Zm0 7L16.575 4H12.4L2 14.425ZM3.4 20H22V4h-2.6Z"/>`),
		g.Group(children),
	)
}