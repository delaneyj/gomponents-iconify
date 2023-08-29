package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassValidLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18h18V6H3v12ZM1 5a1 1 0 0 1 1-1h20a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V5Zm8 5a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm2 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-2.998 6c-.967 0-1.84.39-2.475 1.025l-1.414-1.414A5.486 5.486 0 0 1 8.002 14a5.49 5.49 0 0 1 3.889 1.61l-1.414 1.415A3.486 3.486 0 0 0 8.002 16Zm8.205-1.293l4-4l-1.414-1.414l-3.293 3.293l-1.793-1.793l-1.414 1.414l2.5 2.5l.707.707l.707-.707Z"/>`),
		g.Group(children),
	)
}