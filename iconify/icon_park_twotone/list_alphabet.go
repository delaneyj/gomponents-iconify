package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListAlphabet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTListAlphabet0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 9h22M20 19h22M20 29h22M20 39h22M6 29h6L6 39h6"/><path fill="#555" d="M11 9H7l-.7 7h5.4L11 9Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m6 19l.3-3m5.7 3l-.3-3m0 0L11 9H7l-.7 7m5.4 0H6.3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTListAlphabet0)"/>`),
		g.Group(children),
	)
}