package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroBanknote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M20 21.473a4.5 4.5 0 1 0 0-8.945v8.944ZM8 16.5v1.01H6.5a.5.5 0 0 0 0 1h1.504c.09 1.707 1.51 2.99 3.216 2.99h1.28a.5.5 0 0 0 0-1h-1.28c-1.18 0-2.126-.863-2.213-1.99H11.5a.5.5 0 1 0 0-1H9V16.5h2.5a.5.5 0 0 0 0-1H9.015c.13-1.114 1.077-2 2.205-2h1.28a.5.5 0 0 0 0-1h-1.28c-1.697 0-3.077 1.35-3.21 3H6.5a.5.5 0 0 0 0 1H8Z"/><path d="M1.5 10a3 3 0 0 1 3-3h23a3 3 0 0 1 3 3v18a2.5 2.5 0 0 1-2.5 2.5h-8V25h7.5a1 1 0 0 0 1-1V10a1 1 0 0 0-1-1H19v1h7.5a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H19v6.5h-4V24H5.5a1 1 0 0 1-1-1V11a1 1 0 0 1 1-1H15V9H4.5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1H14v5.5H4A2.5 2.5 0 0 1 1.5 28V10Zm25 1H19v12h7.5V11Zm-21 0v12H15V11H5.5Z"/></g>`),
		g.Group(children),
	)
}