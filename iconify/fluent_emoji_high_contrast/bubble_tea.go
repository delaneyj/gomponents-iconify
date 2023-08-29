package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleTea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M18.65 23.62a1.31 1.31 0 1 0 0-2.62a1.31 1.31 0 0 0 0 2.62Zm-3.99-1.31a1.31 1.31 0 1 1-2.62 0a1.31 1.31 0 0 1 2.62 0Zm2.91 3.3a1.31 1.31 0 1 1-2.62 0a1.31 1.31 0 0 1 2.62 0Zm-5.26 1.31a1.31 1.31 0 1 0 0-2.62a1.31 1.31 0 0 0 0 2.62Z"/><path d="M17 2v5h8a1 1 0 1 1 0 2l-2 20c-.097.605-.5 1-1 1H10c-.5 0-.924-.437-1-1L7 9a1 1 0 0 1 0-2h8V2h2Zm-7.765 9.25L10.91 28h10.18l.16-1.592a1.31 1.31 0 1 1 .137-1.374l1.378-13.784H9.235Z"/></g>`),
		g.Group(children),
	)
}