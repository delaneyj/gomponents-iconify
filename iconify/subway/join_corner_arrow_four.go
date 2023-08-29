package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinCornerArrowFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M511.9 325.8L442.1 256v139.6L116.4 69.9H256L186.2.1L0 0l.1 186.2L69.9 256V116.4l325.7 325.7H256l69.8 69.8l186.2.1z"/>`),
		g.Group(children),
	)
}