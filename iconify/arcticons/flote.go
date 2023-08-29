package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.217 33.758c3.637-3.928 4.365-18.67 0-24.538c9.117 3.83 16.973 13.87 18.186 22.792L7.948 35.407c6.886-8.39 10.135-14.403 12.269-26.187Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.217 9.22V6.795l-6.838 4.267l6.747-1.347M5.547 38.56L42.5 34.293s-.388 6.837-6.79 6.837c-7.08 0-6.934-4.753-6.934-4.753s-1.115 4.223-5.092 4.828H6.614a2.641 2.641 0 0 1-1.067-2.645Z"/>`),
		g.Group(children),
	)
}