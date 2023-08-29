package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PigNose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#E6AAAA" d="M35 18c0 6.075-4.925 11-11 11H12C5.925 29 1 24.075 1 18S5.925 7 12 7h12c6.075 0 11 4.925 11 11z"/><ellipse cx="10" cy="18" fill="#662113" rx="4" ry="6"/><ellipse cx="26" cy="18" fill="#662113" rx="4" ry="6"/>`),
		g.Group(children),
	)
}