package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseCircleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m269 131l30 30l-56 55l56 55l-30 30l-56-55l-55 55l-30-30l55-55l-55-55l30-30l55 55zM213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zm0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50z"/>`),
		g.Group(children),
	)
}