package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAirplaneWindow0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M10 18c0-7.732 6.268-14 14-14s14 6.268 14 14v12c0 7.732-6.268 14-14 14s-14-6.268-14-14V18Z"/><path stroke="#000" d="M10 17h28m-16-6h4M10 26s5.4-.6 7 1c1.6 1.6 1 3.369 1 3.369c3 0 6 .158 6 3.631c0 2.5-4 4-6.5 2.369C17.5 38.5 17 40 14 40m24-17s-3.5-2-5-1s-1 3-1 3c-1.5-1-4 0-4 2.5s2.5 3.5 5 2.5c1 3 3.5 3 5 2"/><path stroke="#fff" d="M25 44v0c7.18 0 13-5.82 13-13V18c0-7.732-6.268-14-14-14v0m-1 40v0c-7.18 0-13-5.82-13-13V18c0-7.732 6.268-14 14-14v0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAirplaneWindow0)"/>`),
		g.Group(children),
	)
}