package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16.07 13c-9.817.35-11.394 8.006-10.956 11.791c-1.753 3.145-1.218 3.93.096 5.24c6.574 6.115 18.843 5.678 24.54 3.495c5.96-2.446 8.999-7.051 9.29-9.526c5.26-3.494 5.366-9.399 4.636-11c-.78 1.31-3.029 2.272-4.635 3c-1.753.35-4.275-.962-5.005-2.127c-.502 2.627 0 4.627 1.314 5.678c2.288 1.747 1.125 3.512.687 3.949c-1.896 1.89-5.506.99-7.26-1.766C24.053 14.31 18.99 13 16.07 13Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 28c2.833.532 8.4 1.554 12-1"/><circle cx="12" cy="20" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}