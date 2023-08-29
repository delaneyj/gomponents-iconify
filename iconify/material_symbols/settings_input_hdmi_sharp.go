package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsInputHdmiSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22v-3l-3-6V7h1V2h12v5h1v6l-3 6v3H8ZM8 7h2V5h1v2h2V5h1v2h2V4H8v3Z"/>`),
		g.Group(children),
	)
}