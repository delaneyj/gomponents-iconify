package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareRoot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M12.08 34.39a1.64 1.64 0 0 1-1.565-1.131l-4.013-12.1l-2.724 1.187a.83.83 0 0 1-1.084-.424a.825.825 0 0 1 .424-1.085l2.719-1.189a1.66 1.66 0 0 1 1.332.003c.423.189.75.55.896.99l4.013 12.099l4.78-25.91a1.652 1.652 0 0 1 1.594-1.222h18.1a.824.824 0 0 1 0 1.646h-18.1l-4.776 25.913a1.641 1.641 0 0 1-1.546 1.221c-.018.002-.034.002-.05.002z"/>`),
		g.Group(children),
	)
}