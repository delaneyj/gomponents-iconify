package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectionOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9 8.25a.75.75 0 0 0 0 1.5h2.25V15a.75.75 0 0 0 1.5 0V9.75H15a.75.75 0 0 0 0-1.5H9Z"/><path fill-rule="evenodd" d="M3.25 6.646A2.751 2.751 0 1 1 6.646 3.25h10.707a2.751 2.751 0 1 1 3.397 3.396v10.707a2.751 2.751 0 1 1-3.396 3.397H6.646a2.751 2.751 0 1 1-3.396-3.396V6.646ZM4 2.75a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5Zm.75 14.604V6.646A2.756 2.756 0 0 0 6.646 4.75h10.707c.26.916.981 1.637 1.897 1.896v10.707a2.756 2.756 0 0 0-1.896 1.897H6.646a2.756 2.756 0 0 0-1.896-1.896ZM4 18.75a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5ZM21.25 4a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm-2.5 16a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}