package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkPenLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.243 4.511L8.505 11.25l-.707 2.12l-1.04 1.042l2.828 2.828l1.04-1.04l2.122-.708l6.737-6.737l-4.242-4.243Zm6.364 3.536a1 1 0 0 1 0 1.414l-7.778 7.778l-2.122.707l-1.414 1.415a1 1 0 0 1-1.414 0l-4.243-4.243a1 1 0 0 1 0-1.414L6.05 12.29l.707-2.122l7.779-7.778a1 1 0 0 1 1.414 0l5.657 5.657Zm-6.364-.707l1.414 1.414l-4.95 4.95l-1.414-1.414l4.95-4.95Zm-10.96 9.546l2.828 2.828l-1.414 1.414l-4.243-1.414l2.829-2.828Z"/>`),
		g.Group(children),
	)
}