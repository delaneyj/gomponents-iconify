package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayerMe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.974 0C9.38 0 4.031 5.349 4.031 11.943v.484C4.375 21.104 13.442 32 13.442 32V17.604c-2.182-.995-3.677-3.156-3.677-5.667c.005-5.063 5.729-8 9.839-5.047c4.109 2.948 3.161 9.313-1.635 10.932l-.042.016v5.901c5.729-.948 10.042-5.87 10.042-11.792c0-6.599-5.349-11.948-11.948-11.948z"/>`),
		g.Group(children),
	)
}