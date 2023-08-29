package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.707 3.293c-.39.39-.369 1.021-.138 1.523a2.83 2.83 0 0 1-3.753 3.753c-.502-.23-1.133-.252-1.523.138l-.586.586a1 1 0 0 0 0 1.414l10.586 10.586a1 1 0 0 0 1.414 0l.586-.586c.39-.39.369-1.021.138-1.523a2.83 2.83 0 0 1 3.753-3.753c.502.23 1.133.252 1.523-.138l.586-.586a1 1 0 0 0 0-1.414L10.707 2.707a1 1 0 0 0-1.414 0l-.586.586Z"/>`),
		g.Group(children),
	)
}