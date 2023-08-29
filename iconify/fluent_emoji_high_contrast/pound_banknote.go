package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoundBanknote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M20 21.473a4.5 4.5 0 1 0 0-8.945v8.944ZM6 17.5a.5.5 0 0 1 .5-.5h1v-2.214A2.245 2.245 0 0 1 9.75 12.5a2.067 2.067 0 0 1 2.25 2a.5.5 0 0 1-1 0c-.005-.237-.112-1-1.25-1a1.256 1.256 0 0 0-1.25 1.29V17H10a.5.5 0 0 1 0 1H8.5v2.29a1.613 1.613 0 0 1-.118.71H12.5a.5.5 0 0 1 0 1h-6a.5.5 0 0 1 0-1h.41c.537 0 .588-.546.59-.714V18h-1a.5.5 0 0 1-.5-.5Z"/><path d="M1.5 10a3 3 0 0 1 3-3h23a3 3 0 0 1 3 3v18a2.5 2.5 0 0 1-2.5 2.5h-8V25h7.5a1 1 0 0 0 1-1V10a1 1 0 0 0-1-1H19v1h7.5a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H19v6.5h-4V24H5.5a1 1 0 0 1-1-1V11a1 1 0 0 1 1-1H15V9H4.5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1H14v5.5H4A2.5 2.5 0 0 1 1.5 28V10Zm25 1H19v12h7.5V11Zm-21 0v12H15V11H5.5Z"/></g>`),
		g.Group(children),
	)
}