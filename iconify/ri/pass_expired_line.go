package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassExpiredLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18h18V6H3v12ZM1 5a1 1 0 0 1 1-1h20a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V5Zm8 5a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm2 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-2.998 6c-.967 0-1.84.39-2.475 1.025l-1.414-1.414A5.486 5.486 0 0 1 8.002 14a5.49 5.49 0 0 1 3.889 1.61l-1.414 1.415A3.486 3.486 0 0 0 8.002 16Zm9.79-7.207L16 10.586l-1.793-1.793l-1.414 1.414L14.586 12l-1.793 1.793l1.414 1.414L16 13.414l1.793 1.793l1.414-1.414L17.414 12l1.793-1.793l-1.414-1.414Z"/>`),
		g.Group(children),
	)
}