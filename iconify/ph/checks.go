package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m149.61 85.71l-89.6 88a8 8 0 0 1-11.22 0L10.39 136a8 8 0 1 1 11.22-11.41l32.79 32.2l84-82.5a8 8 0 1 1 11.22 11.42Zm96.1-11.32a8 8 0 0 0-11.32-.1l-84 82.5l-18.83-18.5a8 8 0 0 0-11.21 11.42l24.43 24a8 8 0 0 0 11.22 0l89.6-88a8 8 0 0 0 .11-11.32Z"/>`),
		g.Group(children),
	)
}