package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compression(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCompression0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M8.92 8.714C8.495 7.39 9.476 6 10.867 6h26.266c1.391 0 2.372 1.39 1.947 2.714C37.9 12.4 36 19.09 36 24s1.9 11.6 3.08 15.286c.425 1.325-.556 2.714-1.947 2.714H10.867c-1.391 0-2.372-1.39-1.947-2.714C10.1 35.6 12 28.91 12 24S10.1 12.4 8.92 8.714Z"/><path stroke-linejoin="round" d="M4 15c1.5 2 2 6 2 9s-.5 7-2 9"/><path d="M18 15.5h12M18 24h12m-12 8h4"/><path stroke-linejoin="round" d="M44 15c-1.5 2-2 6-2 9s.5 6 2 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCompression0)"/>`),
		g.Group(children),
	)
}