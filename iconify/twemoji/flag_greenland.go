package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagGreenland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#D00C33" d="M0 27a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-9H0v9z"/><path fill="#EEE" d="M32 5H4a4 4 0 0 0-4 4v9h36V9a4 4 0 0 0-4-4z"/><circle cx="13.5" cy="18" r="8" fill="#EEE"/><path fill="#D00C33" d="M13.5 10a8 8 0 0 0-8 8h16a8 8 0 0 0-8-8z"/>`),
		g.Group(children),
	)
}