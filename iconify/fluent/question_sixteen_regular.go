package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a3 3 0 0 0-3 3a.5.5 0 0 0 1 0a2 2 0 0 1 4 0c0 .676-.14 1.08-.315 1.36c-.183.293-.428.494-.747.75l-.02.015c-.302.242-.672.537-.951.985C7.673 8.58 7.5 9.176 7.5 10v.5a.5.5 0 0 0 1 0V10c0-.676.14-1.08.315-1.36c.183-.293.428-.494.747-.75l.02-.015c.302-.242.672-.537.951-.985C10.827 6.42 11 5.824 11 5a3 3 0 0 0-3-3Zm0 12a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Z"/>`),
		g.Group(children),
	)
}