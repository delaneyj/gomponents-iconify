package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M18.75 11a7 7 0 1 1-14 0a7 7 0 0 1 14 0Z" opacity=".2"/><path fill-rule="evenodd" d="M10 16a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0 1a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10 6.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0V7a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13.5 10a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1 0-1h3a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}