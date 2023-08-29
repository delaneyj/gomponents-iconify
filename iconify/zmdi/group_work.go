package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zM128 333q22 0 37.5-15.5T181 280t-15.5-37.5T128 227t-37.5 15.5T75 280t15.5 37.5T128 333zm32-202q0 22 15.5 37.5T213 184t38-15.5t16-37.5t-16-38t-38-16t-37.5 16t-15.5 38zm139 202q22 0 37.5-15.5T352 280t-15.5-37.5T299 227t-38 15.5t-16 37.5t16 37.5t38 15.5z"/>`),
		g.Group(children),
	)
}