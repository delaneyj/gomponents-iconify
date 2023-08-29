package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeutralOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M24 40c8.837 0 16-7.163 16-16S32.837 8 24 8S8 15.163 8 24s7.163 16 16 16Zm0 2c9.941 0 18-8.059 18-18S33.941 6 24 6S6 14.059 6 24s8.059 18 18 18Z"/><path d="M30.5 20a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm-2.5.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0ZM17.5 20a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm-2.5.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0ZM16 31a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H17a1 1 0 0 1-1-1Z"/></g>`),
		g.Group(children),
	)
}