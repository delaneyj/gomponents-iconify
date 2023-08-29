package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tongue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zM8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13zM4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm2 4v1h-1v1.5a1.5 1.5 0 0 1-3 0V10H4V9h8z"/>`),
		g.Group(children),
	)
}