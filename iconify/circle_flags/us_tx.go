package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsTx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsUsTx0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsTx0)"><path fill="#0052b4" d="M0 0h167l64 256l-64 256H0Z"/><path fill="#eee" d="m43.5 317l104-75h-128l104 75l-40-122zM167 0h345v256l-173 64l-172-64Z"/><path fill="#d80027" d="M167 256h345v256H167z"/></g>`),
		g.Group(children),
	)
}