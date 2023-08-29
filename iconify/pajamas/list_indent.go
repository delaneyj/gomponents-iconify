package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListIndent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.75 3a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5H4.75Zm2 4.25a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5h-8.5ZM.22 5.72a.75.75 0 0 1 1.06 0l1.75 1.75l.53.53l-.53.53l-1.75 1.75A.75.75 0 1 1 .22 9.22L1.44 8L.22 6.78a.75.75 0 0 1 0-1.06Zm4.53 5.78a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5H4.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}