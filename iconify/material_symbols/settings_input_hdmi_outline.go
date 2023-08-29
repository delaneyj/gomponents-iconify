package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsInputHdmiOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22v-3l-3-6V7h1V4q0-.825.588-1.413T8 2h8q.825 0 1.413.588T18 4v3h1v6l-3 6v3H8ZM8 7h2V5h1v2h2V5h1v2h2V4H8v3Zm2 13h4v-1.5l3-6V9H7v3.5l3 6V20Zm2-5.5Z"/>`),
		g.Group(children),
	)
}