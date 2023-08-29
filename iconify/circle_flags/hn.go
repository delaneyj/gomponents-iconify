package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsHn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsHn0)"><path fill="#338af3" d="M0 0h512v144.7l-40.5 112.6l40.5 110V512H0V367.3l42.2-114L0 144.7z"/><path fill="#eee" d="M0 144.7h512v222.6H0z"/><path fill="#338af3" d="m157.5 167l8.3 25.5h26.9L171 208.2l8.2 25.5l-21.7-15.7l-21.7 15.7l8.3-25.5l-21.7-15.7h26.9zm0 111.3l8.3 25.5h26.9L171 319.5l8.2 25.5l-21.7-15.7l-21.7 15.7l8.3-25.5l-21.7-15.7h26.9zm197-111.3l8.2 25.5h26.9l-21.7 15.7l8.3 25.5l-21.7-15.7l-21.7 15.7l8.2-25.5l-21.7-15.7h26.9zm0 111.3l8.2 25.5h26.9l-21.7 15.7l8.3 25.5l-21.7-15.7l-21.7 15.7l8.2-25.5l-21.7-15.7h26.9zM256 222.6l8.3 25.5H291L269.4 264l8.3 25.5l-21.7-15.8l-21.7 15.8l8.3-25.5l-21.7-15.8h26.8z"/></g>`),
		g.Group(children),
	)
}