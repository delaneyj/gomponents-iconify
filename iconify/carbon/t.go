package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func T(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M12 11h3v12h2V11h3V9h-8v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}