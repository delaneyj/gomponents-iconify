package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsSs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSs0)"><path fill="#eee" d="M74.1 115L512 156.9v31.5l-42.4 70.3l42.4 64.2v31.5L74.1 386.8z"/><path fill="#333" d="M0 0h512v156.8H50z"/><path fill="#a2001d" d="M150.6 188.3H512v134.5H150.6z"/><path fill="#496e2d" d="M50 354.3h462V512H0z"/><path fill="#0052b4" d="M0 0v512l256-256z"/><path fill="#ffda44" d="m83.4 192.4l31.2 43.6l51.2-16.3l-31.9 43.2l31.3 43.6l-51-16.9l-31.7 43.2l.3-53.7L32 262.2L83 246z"/></g>`),
		g.Group(children),
	)
}