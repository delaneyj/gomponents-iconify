package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadialBlurBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="12" r="10" opacity=".5"/><path d="M15.5 9.75a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm-4.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0ZM5.75 9a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm4.75-3.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm4.5 0a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm4 4a.75.75 0 1 0-1.5 0a.75.75 0 0 0 1.5 0Zm0 4.5a.75.75 0 1 0-1.5 0a.75.75 0 0 0 1.5 0Zm-3.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0ZM9.75 15.5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm-4-2a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm4.75 4.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm3.75.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Z"/></g>`),
		g.Group(children),
	)
}