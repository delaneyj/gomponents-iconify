package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyLightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M20.517 14.64v112.715l44.386 28.033l-44.97 78.26l96.948 49.056l-9.343 32.705l78.843 35.04l-32.122 96.947l174.623 46.14l-112.714-69.5l22.778-50.226l257.55 113.885l-282.08-203.82l29.2-90.525l47.89 52.562l83.515-40.88l-44.386 87.016l164.693 108.628l-106.292-126.148l22.778-40.296l74.754 56.648l-84.68-151.26l-71.837 27.45L227.263 14.64H20.52zm77.09 161.19l42.633 26.865l-14.6 50.81l-52.56-34.458l24.528-43.217z"/>`),
		g.Group(children),
	)
}