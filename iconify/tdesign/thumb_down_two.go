package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDownTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.446 22.236l-1.716-.572a3.7 3.7 0 0 1-2.53-3.51V15.7H5.332a3 3 0 0 1-2.965-3.456l1.184-7.7A3 3 0 0 1 6.516 2H22v11.9h-4.85l-3.704 8.336ZM17.5 11.9H20V4h-2.5v7.9Zm-2-7.9H6.517a1 1 0 0 0-.988.848l-1.185 7.7a1 1 0 0 0 .989 1.152H11.2v4.454a1.7 1.7 0 0 0 1.154 1.61l3.146-7.076V4Z"/>`),
		g.Group(children),
	)
}