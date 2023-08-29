package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarXLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 34h-26V24a6 6 0 0 0-12 0v10H86V24a6 6 0 0 0-12 0v10H48a14 14 0 0 0-14 14v160a14 14 0 0 0 14 14h160a14 14 0 0 0 14-14V48a14 14 0 0 0-14-14ZM48 46h26v10a6 6 0 0 0 12 0V46h84v10a6 6 0 0 0 12 0V46h26a2 2 0 0 1 2 2v34H46V48a2 2 0 0 1 2-2Zm160 164H48a2 2 0 0 1-2-2V94h164v114a2 2 0 0 1-2 2Zm-51.76-77.76L136.48 152l19.76 19.76a6 6 0 1 1-8.48 8.48L128 160.48l-19.76 19.76a6 6 0 0 1-8.48-8.48L119.52 152l-19.76-19.76a6 6 0 1 1 8.48-8.48L128 143.52l19.76-19.76a6 6 0 1 1 8.48 8.48Z"/>`),
		g.Group(children),
	)
}