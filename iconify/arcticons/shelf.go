package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shelf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.7 4.5h-2.3c-1.1 0-2 .9-2 2v35c0 1.1.9 2 2 2h2.3m0-39v39h24.9c1.1 0 2-.9 2-2v-35c0-1.1-.9-2-2-2H12.7Z"/><path d="M31.066 5.09v11.236l2.512-2.512m2.538-8.705v11.236l-2.512-2.511m5.804 25.528c-.83.143-1.793-.048-2.694-.96c-3.329-3.362-6.097-11.218-4.022-11.91c.66-.185 1.247.68 1.494 2.286c.436 2.824.084 7.939-.568 11.314c-.172.892-.352 1.962-1 2.935a7.319 7.319 0 0 1-.171.247m.13.183l6.87-3.966"/></g>`),
		g.Group(children),
	)
}