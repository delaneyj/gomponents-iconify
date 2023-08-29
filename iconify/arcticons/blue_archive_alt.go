package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlueArchiveAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.736 4.516L9.79 43.484M8.249 10.253L39.75 29.108m-3.132-2.498c-2.248 3.94-9.858 3.83-16.998-.246c0 0 0 0 0 0c-7.14-4.075-11.105-10.572-8.857-14.51v-.001c2.249-3.939 9.859-3.828 16.998.247c7.14 4.075 11.104 10.571 8.857 14.51Z"/>`),
		g.Group(children),
	)
}