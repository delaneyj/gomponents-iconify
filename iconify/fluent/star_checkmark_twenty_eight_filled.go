package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarCheckmarkTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.302 3.908c-.532-1.078-2.069-1.078-2.6 0L10.01 9.36l-6.017.875c-1.19.173-1.664 1.634-.804 2.473l4.355 4.244l-1.028 5.993c-.203 1.185 1.04 2.088 2.104 1.529l4.547-2.391a7.5 7.5 0 0 1 10.587-8.342l1.06-1.033c.86-.839.385-2.3-.804-2.473l-6.017-.875l-2.691-5.452ZM27 20.5a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0Zm-2.646-2.854a.5.5 0 0 0-.708 0L19 22.293l-1.646-1.647a.5.5 0 0 0-.708.708l2 2a.5.5 0 0 0 .708 0l5-5a.5.5 0 0 0 0-.708Z"/>`),
		g.Group(children),
	)
}