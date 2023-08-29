package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunAbove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1.77 1.01L1 1.42v12l.78.42l9-6v-.83l-9.01-6zM2 12.49V2.36l7.6 5.07L2 12.49zM12.15 8h.71l2.5 2.5l-.71.71L13 9.56V15h-1V9.55l-1.65 1.65l-.7-.7l2.5-2.5z"/>`),
		g.Group(children),
	)
}