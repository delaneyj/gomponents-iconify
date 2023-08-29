package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBotswana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#EEE" d="M0 13h36v10H0z"/><path fill="#75AADB" d="M32 5H4a4 4 0 0 0-4 4v5h36V9a4 4 0 0 0-4-4zM0 27a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-5H0v5z"/><path fill="#141414" d="M0 16h36v4H0z"/>`),
		g.Group(children),
	)
}