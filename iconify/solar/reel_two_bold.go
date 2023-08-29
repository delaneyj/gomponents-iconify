package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReelTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 14.25a2.25 2.25 0 1 1 0-4.5a2.25 2.25 0 0 1 0 4.5Z"/><path fill-rule="evenodd" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10a9.993 9.993 0 0 1-4.73 8.5h3.98a.75.75 0 0 1 0 1.5H12Zm1-16.5a1 1 0 1 0-2 0a1 1 0 0 0 2 0ZM6.5 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm13 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM12 17.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2ZM8.25 12a3.75 3.75 0 1 0 7.5 0a3.75 3.75 0 0 0-7.5 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}