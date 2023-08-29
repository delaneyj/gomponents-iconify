package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PendingFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2ZM8 18a2 2 0 1 1 2-2a2 2 0 0 1-2 2Zm8 0a2 2 0 1 1 2-2a2 2 0 0 1-2 2Zm8 0a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/><path fill="none" d="M10 16a2 2 0 1 1-2-2a2 2 0 0 1 2 2Zm6-2a2 2 0 1 0 2 2a2 2 0 0 0-2-2Zm8 0a2 2 0 1 0 2 2a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}