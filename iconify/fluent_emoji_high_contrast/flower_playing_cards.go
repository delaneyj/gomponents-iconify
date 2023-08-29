package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerPlayingCards(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 7a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12.114A11.674 11.674 0 0 0 22.37 19c-1.409 0-2.752.25-3.98.706C17.136 16.948 14.076 15 10.5 15c-.87 0-1.71.115-2.5.33V7Zm12.5 3.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Z"/><path d="M5 4a2 2 0 0 1 2-2h18a2 2 0 0 1 2 2v24a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V4Zm2 2a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v20a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2V6Z"/></g>`),
		g.Group(children),
	)
}