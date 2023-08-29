package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.5 15.5L15.2 9.552v5.6l-9.7-5.6v11.895l9.7-5.6v5.6z"/>`),
		g.Group(children),
	)
}