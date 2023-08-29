package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMonitor0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 10h32v28H4z"/><path stroke-linecap="round" d="m44 14l-8 6.75v6.5L44 34V14Z" clip-rule="evenodd"/><path stroke-linecap="round" d="m17 19l6 5l-6 5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMonitor0)"/>`),
		g.Group(children),
	)
}