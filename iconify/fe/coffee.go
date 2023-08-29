package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCoffee0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCoffee1" fill="currentColor"><path id="feCoffee2" d="M4 17h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2ZM17 7h1a2 2 0 1 1 0 4h-1c-.033 0-.067 0-.1-.002A5.002 5.002 0 0 1 12 15h-2a5 5 0 0 1-5-5V7a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2ZM7 7v3a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3V7H7Z"/></g></g>`),
		g.Group(children),
	)
}