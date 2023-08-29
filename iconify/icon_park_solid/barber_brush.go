package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarberBrush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBarberBrush0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M29.796 24H16.204s-3.986 7.708 2.548 10.833c4.183 2.5-2.548 9.167-2.548 9.167h13.592s-6.73-7.292-2.548-9.167C33.782 31.708 29.796 24 29.796 24Z"/><path d="m37 10l-7 14H16L9 10s3-6 14-6s14 6 14 6ZM25 24l2-12m-6 12l-2-12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBarberBrush0)"/>`),
		g.Group(children),
	)
}