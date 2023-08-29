package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerPathOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.25 5A3.75 3.75 0 0 1 5 1.25h14A3.75 3.75 0 0 1 21.25 8A3.75 3.75 0 0 1 19 14.75h-6.25v1.604c.916.259 1.637.98 1.896 1.896H22a.75.75 0 0 1 0 1.5h-7.354a2.751 2.751 0 0 1-5.292 0H2a.75.75 0 0 1 0-1.5h7.354a2.756 2.756 0 0 1 1.896-1.896V14.75H5A3.75 3.75 0 0 1 2.75 8a3.744 3.744 0 0 1-1.5-3ZM5 7.25a2.25 2.25 0 0 1 0-4.5h14a2.25 2.25 0 0 1 0 4.5H5Zm14 1.5H5a2.25 2.25 0 0 0 0 4.5h14a2.25 2.25 0 0 0 0-4.5ZM12.25 5a.75.75 0 0 1 .75-.75h6a.75.75 0 0 1 0 1.5h-6a.75.75 0 0 1-.75-.75Zm0 6a.75.75 0 0 1 .75-.75h6a.75.75 0 0 1 0 1.5h-6a.75.75 0 0 1-.75-.75ZM12 17.75a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5Z" clip-rule="evenodd"/><path d="M7 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		g.Group(children),
	)
}