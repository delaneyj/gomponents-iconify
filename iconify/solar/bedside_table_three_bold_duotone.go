package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedsideTableThreeBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2 11h20v1c0 3.771 0 5.657-1.172 6.828a3.1 3.1 0 0 1-1.078.697V22a.75.75 0 0 1-1.5 0v-2.129C17.18 20 15.806 20 14 20h-4c-1.806 0-3.18 0-4.25-.129V22a.75.75 0 0 1-1.5 0v-2.475a3.1 3.1 0 0 1-1.078-.697C2 17.657 2 15.771 2 12v-1Zm6.25 4.5a.75.75 0 0 1 .75-.75h6a.75.75 0 0 1 0 1.5H9a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/><path d="M3.172 3.172C2 4.343 2 6.229 2 10v1h20v-1c0-3.771 0-5.657-1.172-6.828C19.657 2 17.771 2 14 2h-4C6.229 2 4.343 2 3.172 3.172Z" opacity=".5"/><path d="M9 5.75a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5H9Z"/></g>`),
		g.Group(children),
	)
}