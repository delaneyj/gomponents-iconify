package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M15.812 11.729a1.49 1.49 0 0 1-.176.252l-4.984 5.98a1.5 1.5 0 0 1-2.304-1.921l4.2-5.04l-4.2-5.04a1.5 1.5 0 1 1 2.304-1.92l5 6a1.5 1.5 0 0 1 .16 1.689Z" opacity=".2"/><path d="M7.116 4.32a.5.5 0 1 1 .768-.64l5 6a.5.5 0 0 1-.768.64l-5-6Z"/><path d="M7.884 16.32a.5.5 0 0 1-.768-.64l5-6a.5.5 0 1 1 .768.64l-5 6Z"/></g>`),
		g.Group(children),
	)
}