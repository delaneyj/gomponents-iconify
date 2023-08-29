package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsCf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCf0)"><path fill="#ffda44" d="m0 378.4l252.9-28.8L512 378.4V512H322.8L256 481l-66.8 31H0z"/><path fill="#6da544" d="m0 256l249.8-28L512 256v122.4H0z"/><path fill="#eee" d="m0 133.6l255.3-28.3L512 133.6V256H0z"/><path fill="#0052b4" d="M0 0h189.2L256 30l66.8-30H512v133.6H0z"/><path fill="#ffda44" d="m137.7 55.7l6.9 21.2H167L148.9 90l6.9 21.3l-18.1-13.1l-18 13.1l6.8-21.3l-18-13h22.3z"/><path fill="#d80027" d="M189.2 0h133.6v512H189.2z"/></g>`),
		g.Group(children),
	)
}