package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerPlayingCards(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M17 5.05h38v62H17z"/><circle cx="30.987" cy="21.354" r="10" fill="#fff"/><path fill="#3f3f3f" d="M55 48.142c-19.5-19.5-38 .025-38 .025V67h38Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 5.05h38v62H17z"/><circle cx="30.987" cy="21.354" r="10"/><path d="M17 48.167s18.5-19.5 38 0"/></g>`),
		g.Group(children),
	)
}