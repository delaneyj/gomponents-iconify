package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fahrenheit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 3h-7a3 3 0 0 0-3 3v14a1 1 0 0 0 2 0v-7h6a1 1 0 0 0 0-2h-6V6a1 1 0 0 1 1-1h7a1 1 0 0 0 0-2Zm-15 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}