package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagGambia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#CC162B" d="M32 5H4a4 4 0 0 0-4 4v4h36V9a4 4 0 0 0-4-4z"/><path fill="#3C762C" d="M0 27a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-4H0v4z"/><path fill="#EEE" d="M0 21h36v2H0zm0-8h36v2H0z"/><path fill="#0D218A" d="M0 15h36v6H0z"/>`),
		g.Group(children),
	)
}