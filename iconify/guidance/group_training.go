package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupTraining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 18v5m2-6v7m15-6v5m-2-6v7m-13-3.5h13m-2 0l-1-4s-1.5-1-3.5-1s-3.5 1-3.5 1l-1 4M23.5 9v5m-2-6v7M13 11.5h8.5m-2 0l-1-4s-1.5-1-3.5-1s-3.5 1-3.5 1l-.25 1m-2.4 5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Zm6-9s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}