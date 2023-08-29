package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardSim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 45v342q0 17-12.5 29.5T299 429H43q-18 0-30.5-12.5T0 387V131L128 3h171q17 0 29.5 12.5T341 45zM107 365v-42H64v42h43zm170 0v-42h-42v42h42zm-170-85v-85H64v85h43zm85 85v-85h-43v85h43zm0-128v-42h-43v42h43zm85 43v-85h-42v85h42z"/>`),
		g.Group(children),
	)
}