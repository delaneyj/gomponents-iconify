package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsQa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsQa0)"><path fill="#eee" d="M0 0h173l61 255.8L173.4 512H0z"/><path fill="#751a46" d="m173 0l-72.7 30.8L176 63l-75.7 32.2l75.7 32.1l-75.7 32.2l75.7 32.1l-75.7 32.1l75.7 32.2l-75.7 32.2l75.7 32.1l-75.7 32.2l75.7 32.1l-75.7 32.2l75.7 32.1l-75.7 32.2l73.1 31H512V0z"/></g>`),
		g.Group(children),
	)
}