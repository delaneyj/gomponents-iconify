package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSeal0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" d="M5 37a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-5Z"/><path d="M5 31a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V31Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M18.763 15.664a1 1 0 0 1 .942-.664h8.59a1 1 0 0 1 .942.664L34 29H14l4.763-13.336Z"/><rect width="18" height="10.8" x="15" y="4" fill="#555" rx="5.4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSeal0)"/>`),
		g.Group(children),
	)
}