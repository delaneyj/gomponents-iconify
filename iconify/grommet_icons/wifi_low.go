package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-width="2"><path d="M12 20a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-4.243-6.243a6 6 0 0 1 8.486 0"/><path stroke-opacity=".2" d="M4.929 10.929c3.905-3.905 10.237-3.905 14.142 0M2.101 8.1c5.467-5.468 14.331-5.468 19.798 0" opacity=".8"/></g>`),
		g.Group(children),
	)
}