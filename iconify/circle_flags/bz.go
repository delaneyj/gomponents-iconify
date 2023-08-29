package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsBz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBz0)"><path fill="#0052b4" d="m0 44.5l257.8-23.7L512 44.5v423l-252.5 26L0 467.5z"/><circle cx="256" cy="256" r="122.4" fill="#eee"/><circle cx="256" cy="256" r="100.2" fill="#6da544"/><circle cx="256" cy="256" r="66.8" fill="#eee"/><path fill="#0052b4" d="M256 239.3L219.8 256v27.8l36.2 22.3l36.2-22.3V256z"/><path fill="#ffda44" d="M219.8 222.6h72.4V256h-72.4z"/><path fill="#a2001d" d="M0 0h512v44.5H0zm0 467.5h512V512H0z"/></g>`),
		g.Group(children),
	)
}