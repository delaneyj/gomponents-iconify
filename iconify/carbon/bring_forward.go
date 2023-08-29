package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BringForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 15h-2v-3a2.002 2.002 0 0 1 2-2h3v2h-3zm3 15h-3a2.002 2.002 0 0 1-2-2v-3h2v3h3zm3-2h4v2h-4zm10 2h-3v-2h3v-3h2v3a2.002 2.002 0 0 1-2 2zM10 18h2v4h-2zm18 0h2v4h-2zm2-3h-2v-3h-3v-2h3a2.002 2.002 0 0 1 2 2zm-12-5h4v2h-4z"/><path fill="currentColor" d="M8 22H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h16a2.002 2.002 0 0 1 2 2v4h-2V4H4v16h4Z"/>`),
		g.Group(children),
	)
}