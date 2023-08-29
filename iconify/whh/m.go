package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func M(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704.5 1024q-26.5 0-45.5-18.5T640 960V263L440 544q-23 32-56 32q-32 0-56-32L128 263v697q0 27-18.5 45.5t-45 18.5t-45.5-18.5T0 960V64q0-26 18.5-45T64 0q17 0 29 6t15.5 10.5T119 32l265 372L649 32q8-11 11.5-15T676 6.5T704 0q26 0 45 19t19 45v896q0 27-18.5 45.5t-45 18.5z"/>`),
		g.Group(children),
	)
}