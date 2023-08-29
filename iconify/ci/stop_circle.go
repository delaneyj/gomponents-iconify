package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0Z"/><path d="M15 13.4v-2.8c0-.56 0-.84-.11-1.054a.998.998 0 0 0-.436-.437C14.24 9 13.96 9 13.4 9h-2.8c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C9 9.76 9 10.04 9 10.6v2.8c0 .56 0 .84.109 1.054c.096.188.249.34.437.437C9.76 15 10.04 15 10.6 15h2.8c.56 0 .84 0 1.054-.11a.997.997 0 0 0 .437-.436C15 14.24 15 13.96 15 13.4Z"/></g>`),
		g.Group(children),
	)
}