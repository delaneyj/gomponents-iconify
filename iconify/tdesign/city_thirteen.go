package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityThirteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 2H6v2h1v4H1v14h22V8h-6V4h1V2Zm-3 2v16h-2v-5h-2v5H9V4h6ZM7 20H3V10h4v10Zm10 0V10h4v10h-4ZM13.004 5.998H11v2.004h2.004V5.998Z"/>`),
		g.Group(children),
	)
}