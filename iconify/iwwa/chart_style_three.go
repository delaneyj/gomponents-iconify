package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartStyleThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M6.827 25.24h8.156v8.739H6.827zm0-19.225h8.156v18.06H6.827zm18.644 0h8.156v14.364h-8.156zm0 15.445h8.156v12.518h-8.156zm-9.299-3.793h8.156v16.312h-8.156zm0-11.667h8.156v10.682h-8.156z"/>`),
		g.Group(children),
	)
}