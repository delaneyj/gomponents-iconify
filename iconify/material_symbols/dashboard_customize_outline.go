package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardCustomizeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h8v8H3V3Zm2 2v4v-4Zm8-2h8v8h-8V3Zm2 2v4v-4ZM3 13h8v8H3v-8Zm2 2v4v-4Zm11-2h2v3h3v2h-3v3h-2v-3h-3v-2h3v-3Zm-1-8v4h4V5h-4ZM5 5v4h4V5H5Zm0 10v4h4v-4H5Z"/>`),
		g.Group(children),
	)
}