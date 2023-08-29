package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListBulleted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="7" cy="9" r="3" fill="currentColor"/><circle cx="7" cy="23" r="3" fill="currentColor"/><path fill="currentColor" d="M16 22h14v2H16zm0-14h14v2H16z"/>`),
		g.Group(children),
	)
}