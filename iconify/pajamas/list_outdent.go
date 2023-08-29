package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListOutdent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.75 3a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5H4.75Zm2 4.25a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5h-8.5ZM3.34 5.72a.75.75 0 0 0-1.06 0L.53 7.47L0 8l.53.53l1.75 1.75a.75.75 0 1 0 1.06-1.06L2.12 8l1.22-1.22a.75.75 0 0 0 0-1.06Zm1.41 5.78a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5H4.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}