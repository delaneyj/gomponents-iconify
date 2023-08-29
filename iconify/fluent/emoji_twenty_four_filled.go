package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 1.998c5.524 0 10.002 4.478 10.002 10.002c0 5.523-4.479 10-10.002 10c-5.524.001-10.002-4.477-10.002-10C1.998 6.476 6.476 1.998 12 1.998ZM8.462 14.783a.75.75 0 0 0-1.179.928A5.991 5.991 0 0 0 12 18.001a5.99 5.99 0 0 0 4.712-2.284a.75.75 0 1 0-1.177-.93A4.491 4.491 0 0 1 12 16.501c-1.398 0-2.69-.64-3.538-1.718ZM9 8.75a1.25 1.25 0 1 0 0 2.499A1.25 1.25 0 0 0 9 8.75Zm6 0a1.25 1.25 0 1 0 0 2.499a1.25 1.25 0 0 0 0-2.499Z"/>`),
		g.Group(children),
	)
}