package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivacyScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11.6L9.6 4H4q-.825 0-1.413.588T2 6v5.6Zm0 7L16.575 4H12.4L2 14.425V18.6ZM3.4 20H20q.825 0 1.413-.588T22 18V6q0-.825-.588-1.413T20 4h-.6l-16 16Z"/>`),
		g.Group(children),
	)
}