package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lkk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#9D01EB"/><path fill="#FFF" fill-rule="nonzero" d="M10.005 26v-3.656L16 16.24l5.976 6.105V26L16 19.894L10.005 26zM5 13.633h8.469L16 16.24H7.531L5 13.633zm22 0l-2.531 2.606H16V5l2.531 2.586v6.047H27z"/></g>`),
		g.Group(children),
	)
}