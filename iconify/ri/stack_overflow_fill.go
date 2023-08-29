package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackOverflowFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.001 20.003V14.67h2v7.333h-16V14.67h2v5.333h12ZM7.6 14.736l.313-1.979l8.837 1.7l-.113 1.586L7.6 14.736Zm1.2-4.531l.732-1.6l7.998 3.732l-.733 1.6l-7.998-3.732Zm2.265-3.932l1.133-1.333l6.798 5.664l-1.133 1.333l-6.798-5.664Zm4.332-4.132l5.265 7.064l-1.4 1.066l-5.264-7.064l1.4-1.066ZM7.333 18.668V16.67h9.33v2h-9.33Z"/>`),
		g.Group(children),
	)
}