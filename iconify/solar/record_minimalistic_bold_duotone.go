package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M22 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-12 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/><path d="M6 16h12a3.985 3.985 0 0 1-2.646-1H8.646c-.705.622-1.632 1-2.646 1Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}