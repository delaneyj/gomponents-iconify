package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weightlifting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 11v5m2-6v7m21-6v5m-2-6v7m-19-3.5h19m-7.5 2l2 8.35m-6-8.35l-2 8.35m9.5-10.35l-1.5-6s-1.5-1-4-1s-4 1-4 1l-1.5 6m5.35-9s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}