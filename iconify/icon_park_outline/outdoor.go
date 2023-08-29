package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outdoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="m4 42l14-32l10 24l4-12l12 20H4Zm33-28a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/>`),
		g.Group(children),
	)
}