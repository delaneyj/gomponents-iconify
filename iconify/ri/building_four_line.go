package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingFourLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 20h2v2H1v-2h2V3a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v17Zm-2 0V4H5v16h14ZM8 11h3v2H8v-2Zm0-4h3v2H8V7Zm0 8h3v2H8v-2Zm5 0h3v2h-3v-2Zm0-4h3v2h-3v-2Zm0-4h3v2h-3V7Z"/>`),
		g.Group(children),
	)
}