package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUpTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.555 1.764l1.715.572a3.7 3.7 0 0 1 2.53 3.51V8.3h3.869a3 3 0 0 1 2.965 3.456l-1.185 7.7A3 3 0 0 1 17.484 22H2V10.1h4.85l3.705-8.336ZM6.5 12.1H4V20h2.5v-7.9Zm2 7.9h8.984a1 1 0 0 0 .988-.848l1.185-7.7a1 1 0 0 0-.988-1.152H12.8V5.846a1.7 1.7 0 0 0-1.155-1.61L8.5 11.312V20Z"/>`),
		g.Group(children),
	)
}