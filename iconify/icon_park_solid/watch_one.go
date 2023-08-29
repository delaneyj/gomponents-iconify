package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWatchOne0"><g fill="none" stroke-width="4"><path stroke="#fff" stroke-linecap="round" d="M19 14V6a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v8m0 20v8a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-8"/><circle cx="24" cy="24" r="11" fill="#fff" stroke="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M21 25h4m0-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWatchOne0)"/>`),
		g.Group(children),
	)
}