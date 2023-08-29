package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConfusedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M17.5 21a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm-2.5.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0Zm15.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm-2.5.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0Z"/><path d="M24 40c8.837 0 16-7.163 16-16S32.837 8 24 8S8 15.163 8 24s7.163 16 16 16Zm0 2c9.941 0 18-8.059 18-18S33.941 6 24 6S6 14.059 6 24s8.059 18 18 18Z"/><path d="m21.32 34.514l-.044.035a1 1 0 0 1-1.26-1.554l.134-.108c1.184-.96 2.411-1.955 4.274-2.459c1.922-.52 4.409-.494 8.144.323a1 1 0 1 1-.428 1.954c-3.592-.786-5.715-.746-7.193-.346c-1.45.392-2.387 1.15-3.627 2.155Z"/></g>`),
		g.Group(children),
	)
}