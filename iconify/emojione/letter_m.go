package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M37.3 17.5H46v29h-5.7V22.1l-5.5 24.4H29l-5.5-24.4v24.4H18v-29h8.8l5.3 22.8l5.2-22.8"/>`),
		g.Group(children),
	)
}