package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M4 20.07C4 14.507 8.508 10 14.07 10h7.86C27.493 10 32 14.508 32 20.07c0 5.56-4.508 10.069-10.07 10.069H16V35S4 32.57 4 20.07Z"/><path d="M24.477 31.867a8.029 8.029 0 0 0 5.579 2.244H33V38s11-1.944 11-11.944A8.056 8.056 0 0 0 35.945 18h-2.21c.173.656.265 1.345.265 2.056c0 5.794-4.08 10.636-9.523 11.811Z"/></g>`),
		g.Group(children),
	)
}