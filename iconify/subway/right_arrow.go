package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M315.1 48.6H196.9l157.6 157.5H0v78.8h354.5L196.9 442.4h118.2L512 245.5z"/>`),
		g.Group(children),
	)
}