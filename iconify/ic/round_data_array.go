package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundDataArray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 5c0 .55.45 1 1 1h2v12h-2c-.55 0-1 .45-1 1s.45 1 1 1h2c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2h-2c-.55 0-1 .45-1 1zM6 20h2c.55 0 1-.45 1-1s-.45-1-1-1H6V6h2c.55 0 1-.45 1-1s-.45-1-1-1H6c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2z"/>`),
		g.Group(children),
	)
}