package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransgenderSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M11.031 6a1 1 0 0 0-1-1H5.97a1 1 0 0 0-1 1v4.063a1 1 0 1 0 2 0V8.541l1.808 1.857l-1.02 1.019a1 1 0 1 0 1.415 1.414l1-1l.878.901A6.002 6.002 0 0 0 15 22.042V24h-1.5a1 1 0 1 0 0 2H15v1a1 1 0 1 0 2 0v-1h1.5a1 1 0 1 0 0-2H17v-1.958a6.002 6.002 0 0 0 3.972-9.277L25 8.562v1.485a1 1 0 1 0 2 0V6a1 1 0 0 0-1-1h-4.047a1 1 0 1 0 0 2h1.774l-4.137 4.317A5.973 5.973 0 0 0 16 10.125a5.972 5.972 0 0 0-3.56 1.17l-.854-.878l1.121-1.12a1 1 0 0 0-1.414-1.415L10.19 8.984L8.26 7h1.772a1 1 0 0 0 1-1ZM16 12.125a4 4 0 1 1 0 8a4 4 0 0 1 0-8Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}