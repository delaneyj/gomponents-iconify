package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagNiger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#EEE" d="M0 13h36v10H0z"/><path fill="#E05206" d="M32 5H4a4 4 0 0 0-4 4v4h36V9a4 4 0 0 0-4-4z"/><circle cx="18" cy="18" r="4" fill="#E05206"/><path fill="#0DB02B" d="M32 31H4a4 4 0 0 1-4-4v-4h36v4a4 4 0 0 1-4 4z"/>`),
		g.Group(children),
	)
}