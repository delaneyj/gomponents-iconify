package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipOVolumeMute0" width="13" height="13" x="30" y="18" maskUnits="userSpaceOnUse" style="mask-type:alpha"><path fill="currentColor" d="M30 18h13v13H30z"/></mask><g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><g stroke-linecap="round" mask="url(#ipOVolumeMute0)"><path d="m40.735 20.286l-8.486 8.485m.001-8.485l8.485 8.485"/></g><path d="M24 6v36c-7 0-12.201-9.16-12.201-9.16H6a2 2 0 0 1-2-2V17.01a2 2 0 0 1 2-2h5.799S17 6 24 6Z"/></g>`),
		g.Group(children),
	)
}