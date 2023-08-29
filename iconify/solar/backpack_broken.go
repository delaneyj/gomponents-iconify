package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackpackBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M21 10.73v5.645a4.519 4.519 0 0 1-3.538 4.411c-3.598.8-7.326.8-10.923 0A4.519 4.519 0 0 1 3 16.376V10.73a6 6 0 0 1 4.618-5.757l.176-.04l.167-.036a19 19 0 0 1 8.078 0l.167.037l.176.04A5.99 5.99 0 0 1 19.632 7M17.5 15.5V17"/><path d="M15.959 4.5A3 3 0 0 0 13 2h-2a3 3 0 0 0-2.959 2.5M3 14a22.148 22.148 0 0 0 5 1.546M21 14a22.154 22.154 0 0 1-9 1.91M10 13h4"/></g>`),
		g.Group(children),
	)
}