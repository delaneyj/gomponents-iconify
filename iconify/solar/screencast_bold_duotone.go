package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreencastBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M14 4h-4C6.229 4 4.343 4 3.172 5.172C2 6.343 2 8.229 2 12c0 3.771 0 5.657 1.172 6.828C4.343 20 6.229 20 10 20h4c3.771 0 5.657 0 6.828-1.172C22 17.657 22 15.771 22 12c0-3.771 0-5.657-1.172-6.828C19.657 4 17.771 4 14 4Z" opacity=".5"/><path d="M2.75 9.5a.75.75 0 0 0 0 1.5A8.25 8.25 0 0 1 11 19.25a.75.75 0 0 0 1.5 0A9.75 9.75 0 0 0 2.75 9.5Z"/><path d="M2.75 12.5a.75.75 0 0 0 0 1.5C5.65 14 8 16.35 8 19.25a.75.75 0 0 0 1.5 0a6.75 6.75 0 0 0-6.75-6.75Z"/><path d="M2.75 15.5a.75.75 0 0 0 0 1.5A2.25 2.25 0 0 1 5 19.25a.75.75 0 0 0 1.5 0a3.75 3.75 0 0 0-3.75-3.75Z"/></g>`),
		g.Group(children),
	)
}