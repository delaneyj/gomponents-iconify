package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScooterBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 132h-.68l-13.38-40.11l-18.56-55.68A12 12 0 0 0 168 28h-32a12 12 0 0 0 0 24h23.35l13.77 41.3l-55 70.7H83.2a40 40 0 1 0-2.55 24H124a12 12 0 0 0 9.47-4.63l48.77-62.7l6.32 19A40 40 0 1 0 212 132ZM44 188a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm168 0a16 16 0 1 1 16-16a16 16 0 0 1-16 16Z"/>`),
		g.Group(children),
	)
}