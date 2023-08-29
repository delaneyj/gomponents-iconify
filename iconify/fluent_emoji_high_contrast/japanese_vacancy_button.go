package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseVacancyButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M16 5a1 1 0 0 0-1 1v1H7a1 1 0 0 0-1 1v3a1 1 0 1 0 2 0V9h5c0 1.416-1.13 5.06-6.254 6.533a1 1 0 1 0 .552 1.922C13.337 15.719 15 11.246 15 9h2v5a2 2 0 0 0 2 2h4.586A2 2 0 0 0 25 15.414l.707-.707a1 1 0 0 0-1.414-1.414l-.707.707H19V9h6v2a1 1 0 1 0 2 0V8a1 1 0 0 0-1-1h-9V6a1 1 0 0 0-1-1ZM8 19a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2h-6v3h8a1 1 0 1 1 0 2H7a1 1 0 1 1 0-2h8v-3H9a1 1 0 0 1-1-1Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}