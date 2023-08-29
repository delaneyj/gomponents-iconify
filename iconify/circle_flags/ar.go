package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsAr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAr0)"><path fill="#338af3" d="M0 0h512v144.7L488 256l24 111.3V512H0V367.3L26 256L0 144.7z"/><path fill="#eee" d="M0 144.7h512v222.6H0z"/><path fill="#ffda44" d="m332.4 256l-31.2 14.7l16.7 30.3l-34-6.5l-4.2 34.3l-23.7-25.2l-23.6 25.2l-4.3-34.3l-34 6.5l16.6-30.3l-31.2-14.7l31.3-14.7L194 211l34 6.5l4.3-34.3l23.6 25.2l23.6-25.2l4.4 34.3l34-6.5l-16.7 30.3z"/></g>`),
		g.Group(children),
	)
}