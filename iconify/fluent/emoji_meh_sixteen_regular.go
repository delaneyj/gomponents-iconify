package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiMehSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 7a.75.75 0 1 1-1.5 0A.75.75 0 0 1 7 7Zm2.75.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5ZM6 9a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1H6Zm2-7a6 6 0 1 1 0 12A6 6 0 0 1 8 2Zm0 1a5 5 0 1 0 0 10A5 5 0 0 0 8 3Z"/>`),
		g.Group(children),
	)
}