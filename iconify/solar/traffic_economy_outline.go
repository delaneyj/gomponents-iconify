package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficEconomyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9.903 22.543A10.75 10.75 0 0 0 22.75 12a.75.75 0 0 0-1.5 0A9.25 9.25 0 1 1 12 2.75a.75.75 0 0 0 0-1.5a10.75 10.75 0 0 0-2.097 21.293Z"/><path d="M12 8.25a.75.75 0 0 1 .75.75v2.25H15a.75.75 0 0 1 0 1.5h-2.25V15a.75.75 0 0 1-1.5 0v-2.25H9a.75.75 0 0 1 0-1.5h2.25V9a.75.75 0 0 1 .75-.75Zm2.687-6.661a.75.75 0 1 0-.374 1.452a9.267 9.267 0 0 1 6.646 6.646a.75.75 0 0 0 1.452-.374a10.768 10.768 0 0 0-7.724-7.724Z"/></g>`),
		g.Group(children),
	)
}