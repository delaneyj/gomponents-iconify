package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimpleCalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.099 7.859a1.92 1.92 0 0 0-1.92 1.92V41.58A1.92 1.92 0 0 0 8.1 43.5h31.8a1.92 1.92 0 0 0 1.92-1.92V9.779a1.92 1.92 0 0 0-1.92-1.92Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.068 33.543h7.497v7.078H9.068zm11.183 0h7.497v7.078h-7.497zm11.184 0h7.497v7.078h-7.497zM9.068 23.335h7.497v7.078H9.068zm11.183 0h7.497v7.078h-7.497zm11.184 0h7.497v7.078h-7.497zm0-9.738h7.497v7.078h-7.497zm-18.618-3.339V4.5m22.366 5.758V4.5"/>`),
		g.Group(children),
	)
}