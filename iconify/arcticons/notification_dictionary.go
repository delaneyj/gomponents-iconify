package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationDictionary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.622v27.166c7.078-3.906 13.973-3.233 19.5 0c5.997-3.298 12.514-3.603 19.5 0V11.622a20.936 20.936 0 0 0-19.5 0c-5.461-1.673-9.745-4.379-19.5 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.684 19.538a17.915 17.915 0 0 1 11.43-.807m-11.43 10.221a17.915 17.915 0 0 1 11.43-.807m-11.43-4.035a17.915 17.915 0 0 1 11.43-.807"/>`),
		g.Group(children),
	)
}