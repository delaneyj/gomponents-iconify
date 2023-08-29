package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 2h12v2h-1v4h2v14H2V8h8V4H9V2Zm3 2v4h6V4h-6Zm8 6H4v10h4v-6h8v6h4V10Zm-6 10v-4h-4v4h4ZM13 4.998h2.004v2.004H13V4.998Z"/>`),
		g.Group(children),
	)
}