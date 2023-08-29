package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19.0099 42H9C7.34315 42 6 40.6569 6 39V9C6 7.34315 7.34315 6 9 6H39C40.6569 6 42 7.34315 42 9V19.0304"/><path d="M42 29.0347V41.0001C42 41.5524 41.5523 42.0001 41 42.0001H29.037"/><path d="M42 29.0347H18"/><path d="M23 23L17 29L23 35"/></g>`),
		g.Group(children),
	)
}