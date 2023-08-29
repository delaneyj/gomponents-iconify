package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsCw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCw0)"><path fill="#0052b4" d="M0 0h512v342.3l-22 34.2l22 32.5v103H0V409l25.4-31L0 342.2z"/><path fill="#eee" d="m175.2 164.2l13.8 42.5h44.7L197.6 233l13.8 42.5l-36.2-26.3l-36.1 26.3l13.8-42.5l-36.2-26.3h44.7zm-76.7-44.5l8.2 25.5h26.9L111.9 161l8.3 25.5l-21.7-15.7l-21.7 15.7L85 161l-21.7-15.7h26.9z"/><path fill="#ffda44" d="M0 342.3h512V409H0z"/></g>`),
		g.Group(children),
	)
}