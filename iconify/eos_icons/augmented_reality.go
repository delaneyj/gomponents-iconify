package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AugmentedReality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 3.311L19 7.653v8.694l-7.5 4.342L4 16.347V7.653l7.5-4.342M11.5 1L2 6.5v11l9.5 5.5l9.5-5.5v-11L11.5 1Z"/><path fill="currentColor" d="M11.5 5a.5.5 0 1 0 .5.5a.5.5 0 0 0-.5-.5Zm0 3.5a.5.5 0 1 0 .5.5a.5.5 0 0 0-.5-.5Zm0 3.5a.5.5 0 1 0 .5.5a.5.5 0 0 0-.5-.5ZM8 13a.5.5 0 1 0 .422.231A.498.498 0 0 0 8 13Zm-2.5 2a.5.5 0 1 0 .421.231a.498.498 0 0 0-.422-.23Zm9.5-2a.5.5 0 1 0 .269.079a.5.5 0 0 0-.269-.08Zm2.535 2a.5.5 0 1 0 .269.079a.5.5 0 0 0-.269-.079ZM3.382 6.225l-1 1.732l7.027 4.057l1-1.732l-7.027-4.057zM12.5 17.016h-2v5.141h2v-5.141z"/>`),
		g.Group(children),
	)
}