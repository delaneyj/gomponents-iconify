package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bytedance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.877 1.468L24 2.534v18.942l-4.123 1.056V1.468zM6.53 10.898l4.115 1.064v8.978L6.53 22.003V10.896zM0 2.572l4.115 1.064v16.736L0 21.428V2.572zm17.455 5.62V19.3l-4.122-1.064V9.257l4.122-1.064z"/>`),
		g.Group(children),
	)
}