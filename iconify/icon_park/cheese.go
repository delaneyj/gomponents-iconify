package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cheese(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M43 20C43 17.8285 24.8921 8.11198 20.134 5.59629C19.4394 5.22904 18.603 5.31194 17.9852 5.79737L5 16"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 17.6522C5 16.3232 6.2688 15.3543 7.55521 15.688C13.9619 17.3498 30.8602 21.3331 40.1615 19.7589C41.5557 19.523 43 20.5369 43 21.951V38.1025C43 39.1662 42.1674 40.0438 41.1051 40.0997L7.10512 41.8892C5.96083 41.9494 5 41.0378 5 39.892V17.6522Z"/><circle cx="12" cy="25" r="2" fill="#fff"/><circle cx="25" cy="27" r="2" fill="#fff"/><circle cx="34" cy="32" r="2" fill="#fff"/><circle cx="18" cy="32" r="2" fill="#fff" stroke="#fff" stroke-width="2"/></g>`),
		g.Group(children),
	)
}