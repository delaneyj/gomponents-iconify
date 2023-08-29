package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SleepTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSleepTwo0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 12v24m40-7v7m0-7H4"/><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M22 16v13h22V19a3 3 0 0 0-3-3H22Z"/><circle cx="13" cy="20" r="3" fill="#555"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSleepTwo0)"/>`),
		g.Group(children),
	)
}