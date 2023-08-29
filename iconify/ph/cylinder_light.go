package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CylinderLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 18c-39.25 0-70 18.45-70 42v136c0 23.55 30.75 42 70 42s70-18.45 70-42V60c0-23.55-30.75-42-70-42Zm0 12c31.44 0 58 13.74 58 30s-26.56 30-58 30s-58-13.74-58-30s26.56-30 58-30Zm0 196c-31.44 0-58-13.74-58-30V83.81C82.48 94.87 103.59 102 128 102s45.52-7.13 58-18.19V196c0 16.26-26.56 30-58 30Z"/>`),
		g.Group(children),
	)
}