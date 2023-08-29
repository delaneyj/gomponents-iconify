package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSwitzerland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#D32D27" d="M31 27a4 4 0 0 1-4 4H9a4 4 0 0 1-4-4V9a4 4 0 0 1 4-4h18a4 4 0 0 1 4 4v18z"/><path fill="#FFF" d="M25 16.063h-5v-5h-4v5h-5V20h5v5.063h4V20h5z"/>`),
		g.Group(children),
	)
}