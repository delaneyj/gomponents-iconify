package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M627.277 0v545.462L1172.723 0v1200L627.262 654.538V1200L27.277 600l600-600z"/>`),
		g.Group(children),
	)
}