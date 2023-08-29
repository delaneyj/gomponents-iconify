package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodTongueSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6 10.5V9h3v1.5a1.5 1.5 0 0 1-3 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM5 6H4V5h1v1Zm6 0h-1V5h1v1ZM4 9h1v1.5a2.5 2.5 0 0 0 5 0V9h1V8H4v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}