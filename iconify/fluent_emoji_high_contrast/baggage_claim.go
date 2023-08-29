package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaggageClaim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M13.25 8.75c0 .084.004.168.011.25H11.25a.25.25 0 0 0-.25.25v11.5c0 .138.112.25.25.25h9.5a.25.25 0 0 0 .25-.25V9.25a.25.25 0 0 0-.25-.25h-2.011a2.75 2.75 0 1 0-5.489-.25Zm4.5 0c0 .085-.006.168-.018.25h-3.464a1.75 1.75 0 1 1 3.482-.25ZM7 10a1 1 0 0 1 1-1h1.75a.25.25 0 0 1 .25.25v11.5a.25.25 0 0 1-.25.25H8a1 1 0 0 1-1-1V10Zm15.25 11a.25.25 0 0 1-.25-.25V9.25a.25.25 0 0 1 .25-.25H24a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-1.75Zm3.25 2a.5.5 0 0 1 0 1H23a1 1 0 0 1-2 0h-2a1 1 0 0 1-2 0h-2a1 1 0 0 1-2 0h-2a1 1 0 0 1-2 0H6.5a.5.5 0 0 1 0-1h19Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}