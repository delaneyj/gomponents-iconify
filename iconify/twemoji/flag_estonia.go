package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagEstonia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#141414" d="M0 14h36v9H0z"/><path fill="#4891D9" d="M32 5H4a4 4 0 0 0-4 4v5h36V9a4 4 0 0 0-4-4z"/><path fill="#EEE" d="M32 31H4a4 4 0 0 1-4-4v-4h36v4a4 4 0 0 1-4 4z"/>`),
		g.Group(children),
	)
}