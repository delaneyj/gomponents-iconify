package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Olive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M14.427 9.994a4.43 4.43 0 1 1-8.86 0a4.43 4.43 0 0 1 8.86 0Z"/><path d="M25.974 6.016C19.868-.079 10.39-.782 4.8 4.796c-5.591 5.582-4.875 15.06 1.23 21.165c6.105 6.105 15.583 6.821 21.174 1.23c5.591-5.59 4.875-15.07-1.23-21.175Zm-19.76.196c4.65-4.641 12.852-4.265 18.346 1.22c5.495 5.494 5.88 13.696 1.23 18.345c-4.65 4.65-12.85 4.265-18.346-1.23c-5.495-5.495-5.878-13.696-1.23-18.335Z"/></g>`),
		g.Group(children),
	)
}