package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zm-.5 384q58 0 105-36L79 111q-36 47-36 105q0 71 50 121t120 50zm135-66q36-47 36-105q0-71-50-121T213 45q-58 0-104 36z"/>`),
		g.Group(children),
	)
}