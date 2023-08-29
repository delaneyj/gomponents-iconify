package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Album(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zm-.5 309q40 0 68-28t28-68t-28-68t-68-28t-68 28t-28 68t28 68t68 28zm.5-117q8.5 0 15 6t6.5 15t-6.5 15t-15 6t-15-6t-6.5-15t6.5-15t15-6z"/>`),
		g.Group(children),
	)
}