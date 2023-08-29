package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectionBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M8.25 9A.75.75 0 0 1 9 8.25h6a.75.75 0 0 1 0 1.5h-2.25V15a.75.75 0 0 1-1.5 0V9.75H9A.75.75 0 0 1 8.25 9Z" clip-rule="evenodd"/><path d="M4 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM22 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-2 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path d="M4.75 5.855a1.994 1.994 0 0 1-1.5 0v12.29a1.993 1.993 0 0 1 1.5 0V5.855ZM5.855 4.75h12.29a1.993 1.993 0 0 1 0-1.5H5.855a1.994 1.994 0 0 1 0 1.5ZM19.25 5.855a1.994 1.994 0 0 0 1.5 0v12.29a1.994 1.994 0 0 0-1.5 0V5.855ZM18.145 19.25H5.855a1.994 1.994 0 0 1 0 1.5h12.29a1.994 1.994 0 0 1 0-1.5Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}