package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagThailand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#A7122D" d="M0 26.518V27a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-.482H0z"/><path fill="#EEE" d="M0 22.181h36v4.485H0z"/><path fill="#292648" d="M0 13.513h36v8.821H0z"/><path fill="#EEE" d="M0 9.181h36v4.485H0z"/><path fill="#A7122D" d="M0 9.333V9a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v.333H0z"/>`),
		g.Group(children),
	)
}