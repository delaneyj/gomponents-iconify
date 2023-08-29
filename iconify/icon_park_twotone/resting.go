package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Resting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTResting0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 26v12m0-32v20m0-22v2M8 40h26l6-6m-25 6v4m17-4v4"/><path fill="#555" d="M24 6C14.059 6 6 14.034 6 23.944V26c0-2.478 2.015-5 4.5-5s4.5 2.522 4.5 5c0-2.478 2.015-5 4.5-5s4.5 2.522 4.5 5c0-2.478 2.015-5 4.5-5s4.5 2.522 4.5 5c0-2.478 2.015-5 4.5-5s4.5 2.522 4.5 5v-2.056C42 14.034 33.941 6 24 6Z"/><path d="M15 26s-1.5-5.5 1-11c2.501-5.5 8-9 8-9m9 20s.5-4.5-2-11s-7-9-7-9m0 19V6"/><path d="M28.5 21c2.485 0 4.5 2.522 4.5 5c0-2.478 2.015-5 4.5-5m-9 0c-2.485 0-4.5 2.522-4.5 5c0-2.478-2.015-5-4.5-5m-9 0c2.485 0 4.5 2.522 4.5 5c0-2.478 2.015-5 4.5-5M15.733 8C18.21 6.722 21.02 6 24 6c2.98 0 5.79.722 8.266 2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTResting0)"/>`),
		g.Group(children),
	)
}