package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M277 462L174 282L334 2h108L282 282l103 180H277zM108 323l80-132l-60-105H26l60 105L6 323h102z"/>`),
		g.Group(children),
	)
}