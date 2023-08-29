package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnstageAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 3.75a.75.75 0 0 1 1.5 0v.943a5 5 0 1 1 .67 7.246a.75.75 0 1 1 .924-1.182A3.5 3.5 0 1 0 2.088 6.5H4.25a.75.75 0 0 1 0 1.5H0V3.75Zm12 0a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75ZM12 8a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 12 8Zm.75 3.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}