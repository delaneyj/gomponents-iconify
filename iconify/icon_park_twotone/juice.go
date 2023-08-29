package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Juice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTJuice0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M15 24h18l-1.8 20H16.8L15 24Z"/><rect width="26" height="6" x="11" y="18" rx="3"/><path fill="#555" d="M24 8c-5.523 0-10 4.477-10 10h20c0-5.523-4.477-10-10-10Z"/><path stroke-linecap="round" d="m28 4l-2 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTJuice0)"/>`),
		g.Group(children),
	)
}