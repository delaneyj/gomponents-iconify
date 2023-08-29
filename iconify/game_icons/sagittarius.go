package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sagittarius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m267.934 459.625l-80.013-80.08l-100.315 100.12l-57.517-57.516l100.25-100.252c-60.47-60.56-77.15-77.326-79.827-80.078l57.52-57.522l79.95 79.952l128.03-128.028C178.14 101.764 209.1 109.4 204.28 108.128L223.96 29.2l203.814 50.813L477.8 283.637l-79.192 19.745l-26.762-107.595l-126.212 126.106l80.02 80.018l-57.72 57.715z"/>`),
		g.Group(children),
	)
}