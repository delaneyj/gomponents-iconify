package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleSpeak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPeopleSpeak0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M19 20a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path d="M33 8s2.25 4.5 0 10m7-14s4.5 8.1 0 18"/><path fill="#555" d="M4 40.8V42h30v-1.2c0-4.48 0-6.72-.872-8.432a8 8 0 0 0-3.496-3.496C27.92 28 25.68 28 21.2 28h-4.4c-4.48 0-6.72 0-8.432.872a8 8 0 0 0-3.496 3.496C4 34.08 4 36.32 4 40.8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPeopleSpeak0)"/>`),
		g.Group(children),
	)
}