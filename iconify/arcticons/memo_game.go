package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoGame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h10v10h-10zm13.5 0h10v10H19zm13.5 0h10v10h-10zM5.5 19h10v10h-10zM19 19h10v10H19zm13.5 0h10v10h-10zm-27 13.5h10v10h-10zm13.5 0h10v10H19zm13.5 0h10v10h-10z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 22h4v4h-4zm27-13.5h4v4h-4z"/>`),
		g.Group(children),
	)
}