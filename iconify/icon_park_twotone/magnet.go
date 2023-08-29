package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMagnet0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 10v16c0 9.941 8.059 18 18 18s18-8.059 18-18V10m-28 0v16c0 5.523 4.477 10 10 10s10-4.477 10-10V10"/><path fill="#555" d="M14 4H6v6h8V4Zm20 0h8v6h-8V4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMagnet0)"/>`),
		g.Group(children),
	)
}