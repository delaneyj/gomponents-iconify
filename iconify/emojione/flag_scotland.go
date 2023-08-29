package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagScotland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="none"><circle cx="32" cy="32" r="30" fill="#0065BD"/><path fill="#EEE" d="M57.972 16.974L38.539 32l19.433 15.026a30.05 30.05 0 0 1-2.046 3.075l-.36.466a30.18 30.18 0 0 1-2.492 2.785L32 37.056L10.926 53.352a30.166 30.166 0 0 1-2.49-2.782l-.364-.471a30.044 30.044 0 0 1-2.044-3.073L25.461 32L6.028 16.974A30.045 30.045 0 0 1 8.074 13.9l.36-.466a30.158 30.158 0 0 1 2.492-2.784L32 26.944l21.074-16.296c.887.876 1.72 1.807 2.494 2.787l.356.461a30.04 30.04 0 0 1 2.048 3.078z"/></g>`),
		g.Group(children),
	)
}