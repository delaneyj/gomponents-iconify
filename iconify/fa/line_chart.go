package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M2048 1408v128H0V0h128v1408h1920zM1920 160v435q0 21-19.5 29.5T1865 617l-121-121l-633 633q-10 10-23 10t-23-10L832 896l-416 416l-192-192l585-585q10-10 23-10t23 10l233 233l464-464l-121-121q-16-16-7.5-35.5T1453 128h435q14 0 23 9t9 23z"/>`),
		g.Group(children),
	)
}