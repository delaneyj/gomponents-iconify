package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zM309 301q22 0 38-15.5t16-37.5t-16-37.5t-38-15.5h-10q0-36-25-61t-61-25q-29 0-52 18.5T131 174l-3-1q-27 0-45.5 19T64 237.5t18.5 45T128 301h181z"/>`),
		g.Group(children),
	)
}