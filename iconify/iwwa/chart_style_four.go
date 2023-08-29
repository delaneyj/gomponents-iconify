package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartStyleFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M6.978 25.24h8.156v8.739H6.978zm0-9.904h8.156v8.739H6.978zm9.323 2.331h8.156v16.312h-8.156zm0-6.991h8.156v5.826h-8.156zm9.322-4.661h8.156v27.964h-8.156z"/>`),
		g.Group(children),
	)
}