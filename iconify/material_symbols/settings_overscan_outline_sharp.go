package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsOverscanOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14v-4l-2 2l2 2Zm5 3l2-2h-4l2 2Zm-2-8h4l-2-2l-2 2Zm7 5l2-2l-2-2v4ZM2 20V4h20v16H2Zm2-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}