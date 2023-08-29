package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quechua(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsQuechua0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsQuechua0)"><path fill="#4a1f63" d="M0 512h512v-70l-256-32L0 442Z"/><path fill="#0052b4" d="M0 442h512v-70l-256-32L0 372Z"/><path fill="#d80027" d="M0 0h512v70l-256 32L0 70Z"/><path fill="#ffda44" d="M0 70h512v70l-256 32L0 140Z"/><path fill="#eee" d="M0 140h512v70l-32 46l32 46v70H0v-70l32-46l-32-46Z"/><path fill="#496e2d" d="M0 210h512v92H0z"/></g>`),
		g.Group(children),
	)
}