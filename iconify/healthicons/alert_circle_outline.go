package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20S35.046 4 24 4Zm0 2C14.059 6 6 14.059 6 24s8.059 18 18 18s18-8.059 18-18S33.941 6 24 6Z" clip-rule="evenodd"/><path d="M26 13a2 2 0 1 0-4 0v14a2 2 0 1 0 4 0V13Zm-2 20a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/></g>`),
		g.Group(children),
	)
}