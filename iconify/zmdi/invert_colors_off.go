package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvertColorsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m377 397l7 8l-27 27l-58-58q-46 38-107 38q-71 0-121-50q-46-46-49.5-112T59 134L0 75l27-27l59 59l30 30l76 76l134 134zm-185-27V267L90 165q-26 34-26 77q0 53 38 90q37 38 90 38zm0-309l-49 48l-30-30l79-79l121 121q38 39 47 92.5T347 313L192 159V61z"/>`),
		g.Group(children),
	)
}