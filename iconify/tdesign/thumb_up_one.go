package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUpOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.555 1.764l1.715.572a3.7 3.7 0 0 1 2.53 3.51V8.3h3.869a3 3 0 0 1 2.965 3.456l-1.185 7.7A3 3 0 0 1 17.484 22H3V10.1h3.85l3.705-8.336Zm1.09 2.472L8.15 12.1H5V20h12.484a1 1 0 0 0 .988-.848l1.185-7.7a1 1 0 0 0-.988-1.152H12.8V5.846a1.7 1.7 0 0 0-1.155-1.61Z"/>`),
		g.Group(children),
	)
}