package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cactus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6.5V14h-6v6H6v-4H0V8.5a2.5 2.5 0 0 1 5 0V11h1V4a4 4 0 0 1 7.988-.314V9H15V6.5a2.5 2.5 0 1 1 5 0zM8 18h4v-6h6V6.5a.5.5 0 1 0-1 0V11h-5V4a2 2 0 1 0-4 0v9H3V8.5a.5.5 0 0 0-1 0V14h6v4zm2-9a1 1 0 0 1 1 1v6a1 1 0 0 1-2 0v-6a1 1 0 0 1 1-1zm0-5a1 1 0 0 1 1 1v2a1 1 0 0 1-2 0V5a1 1 0 0 1 1-1z"/>`),
		g.Group(children),
	)
}