package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Default(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 9v26a3 3 0 0 0 3 3h11.832A9.966 9.966 0 0 1 22 34c0-5.523 4.477-10 10-10c.685 0 1.354.069 2 .2V9a3 3 0 0 0-3-3H11a3 3 0 0 0-3 3Zm5 4a1 1 0 1 0 0 2h16a1 1 0 1 0 0-2H13Zm0 5a1 1 0 1 0 0 2h16a1 1 0 1 0 0-2H13Zm-1 6a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2h-8a1 1 0 0 1-1-1Zm1 4a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2h-5Z"/><path d="M32 42a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm2.567-11.707a1 1 0 0 1 1.414 1.414l-.79.791l-1.415-1.414l.791-.791ZM28 36.86l5.069-5.07l1.414 1.415l-5.069 5.069H28V36.86Z"/></g>`),
		g.Group(children),
	)
}