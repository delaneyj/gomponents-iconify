package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleEditTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2a8 8 0 0 0-.992 15.94c.009-.095.025-.19.05-.287l.374-1.498a3.19 3.19 0 0 1 .84-1.485l4.83-4.83a2.87 2.87 0 0 1 2.85-.72A8.001 8.001 0 0 0 10 2Zm.98 13.377l4.83-4.83a1.87 1.87 0 1 1 2.644 2.646l-4.83 4.829a2.197 2.197 0 0 1-1.02.578l-1.498.374a.89.89 0 0 1-1.079-1.078l.375-1.498a2.18 2.18 0 0 1 .578-1.02Z"/>`),
		g.Group(children),
	)
}