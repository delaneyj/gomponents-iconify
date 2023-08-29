package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kitchen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 22.5H1.5v-20s2-1 5.25-1s5.25 1 5.25 1v20Zm0 0v-10h.25s1.75 1 5 1s5-1 5-1h.25v10H12ZM1.5 9.5s2-1 5.25-1s5.25 1 5.25 1m-8-5v2M4 11v2m14.5.449V9.75a1.25 1.25 0 1 0-2.5 0V10"/>`),
		g.Group(children),
	)
}