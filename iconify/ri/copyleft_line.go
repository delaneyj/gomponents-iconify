package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyleftLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22C6.48 22 2 17.52 2 12S6.48 2 12 2s10 4.48 10 10s-4.48 10-10 10Zm0-2c4.42 0 8-3.58 8-8s-3.58-8-8-8s-8 3.58-8 8s3.58 8 8 8Zm0-3a4.999 4.999 0 0 1-4.288-2.428l1.714-1.029A3 3 0 1 0 12 9a2.998 2.998 0 0 0-2.573 1.456L7.712 9.428A4.999 4.999 0 0 1 17 12c0 2.76-2.24 5-5 5Z"/>`),
		g.Group(children),
	)
}