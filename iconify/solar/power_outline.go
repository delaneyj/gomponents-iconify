package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.75 2a.75.75 0 0 0-1.5 0v4a.75.75 0 0 0 1.5 0V2Z"/><path d="M8.792 4.397a.75.75 0 1 0-.584-1.382A9.753 9.753 0 0 0 2.25 12c0 5.385 4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75a9.752 9.752 0 0 0-5.958-8.985a.75.75 0 1 0-.584 1.382A8.253 8.253 0 0 1 12 20.25A8.25 8.25 0 0 1 8.792 4.397Z"/></g>`),
		g.Group(children),
	)
}