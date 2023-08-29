package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusNeutral(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 12A6 6 0 1 0 6 0a6 6 0 0 0 0 12zM3 4.75A.75.75 0 0 1 3.75 4h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 3 4.75zm0 2.5a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 0 1.5h-3.5A.75.75 0 0 1 3 7.25z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}