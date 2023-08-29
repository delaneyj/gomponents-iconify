package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IcecreamTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSIcecreamTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M15.34 22.5L21 37l3 6l3-6l5.66-14.5"/><path d="M19 32h10"/><path fill="#fff" stroke-linejoin="round" d="M24 3c-6 0-8 6-8 6s-6 2-6 7s5 7 5 7s3.5-2 9-2s9 2 9 2s5-2 5-7s-6-7-6-7s-2-6-8-6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSIcecreamTwo0)"/>`),
		g.Group(children),
	)
}