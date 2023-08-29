package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M17.217 5.333a1 1 0 1 1-1.49 1.334A5 5 0 0 0 7 10h2a1 1 0 0 1 .707 1.707l-3 3a1 1 0 0 1-1.414 0l-3-3A1 1 0 0 1 3 10h2a7 7 0 0 1 12.217-4.667zm4.49 6.96A1 1 0 0 1 21 14h-2v.002a7 7 0 0 1-12.217 4.665a1 1 0 1 1 1.49-1.334A5 5 0 0 0 17 14h-2a1 1 0 0 1-.707-1.707l3-3a1 1 0 0 1 1.414 0l3 3z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}