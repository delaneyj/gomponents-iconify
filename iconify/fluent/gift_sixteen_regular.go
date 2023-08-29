package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 2c.597 0 1.134.262 1.5.677A2 2 0 0 1 10.732 5H12a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h1.268A2 2 0 0 1 6 2ZM5 4a1 1 0 0 0 1 1h1V4a1 1 0 0 0-2 0Zm3 2v2h4V6H8ZM7 6H3v2h4V6ZM4 9v3a1 1 0 0 0 1 1h2V9H4Zm4 4h2a1 1 0 0 0 1-1V9H8v4Zm2-9a1 1 0 0 0-2 0v1h1a1 1 0 0 0 1-1Z"/>`),
		g.Group(children),
	)
}