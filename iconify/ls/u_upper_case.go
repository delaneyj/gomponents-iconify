package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M527 0v491c0 71-28 134-72 181c-48 51-115 82-191 82s-144-31-192-82C28 625 0 561 0 491V0h72v492c1 105 87 191 192 191s190-86 191-190V0h72z"/>`),
		g.Group(children),
	)
}