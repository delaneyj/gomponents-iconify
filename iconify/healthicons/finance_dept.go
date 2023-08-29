package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinanceDept(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Zm8 4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2H17Z" clip-rule="evenodd"/><path d="M17 12h14v9H17v-9Zm4.273 12.636a1.636 1.636 0 1 1-3.273 0a1.636 1.636 0 0 1 3.273 0Zm4.363 0a1.636 1.636 0 1 1-3.272 0a1.636 1.636 0 0 1 3.272 0Zm4.364 0a1.636 1.636 0 1 1-3.273 0a1.636 1.636 0 0 1 3.273 0ZM21.273 29A1.636 1.636 0 1 1 18 29a1.636 1.636 0 0 1 3.273 0Zm4.363 0a1.636 1.636 0 1 1-3.273 0a1.636 1.636 0 0 1 3.273 0ZM30 29a1.636 1.636 0 1 1-3.273 0A1.636 1.636 0 0 1 30 29Zm0 4.364a1.636 1.636 0 1 1-3.273 0a1.636 1.636 0 0 1 3.273 0Zm-12 0c0-.904.733-1.637 1.636-1.637H24A1.636 1.636 0 1 1 24 35h-4.364A1.636 1.636 0 0 1 18 33.364Z"/></g>`),
		g.Group(children),
	)
}