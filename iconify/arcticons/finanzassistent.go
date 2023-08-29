package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Finanzassistent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.6 36.87c0-12.29 16-25.74 17-25.74H20.1c-2 0-15.6 13.45-15.6 25.74Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.67 35.42h9.68c0-8.58 11.16-18 11.9-18H28.43"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.43 34.14h6.69c0-6 7.85-12.67 8.38-12.67h-7.85"/>`),
		g.Group(children),
	)
}