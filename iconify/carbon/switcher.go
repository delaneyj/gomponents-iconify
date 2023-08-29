package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Switcher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 4h4v4h-4zM4 4h4v4H4zm20 0h4v4h-4zM14 14h4v4h-4zM4 14h4v4H4zm20 0h4v4h-4zM14 24h4v4h-4zM4 24h4v4H4zm20 0h4v4h-4z"/>`),
		g.Group(children),
	)
}