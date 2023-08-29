package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CylinderFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 16c-40.37 0-72 19.33-72 44v136c0 24.67 31.63 44 72 44s72-19.33 72-44V60c0-24.67-31.63-44-72-44Zm0 208c-29.83 0-56-13.08-56-28V77.43C82.92 88.5 103.9 96 128 96s45.08-7.5 56-18.57V196c0 14.92-26.17 28-56 28Z"/>`),
		g.Group(children),
	)
}