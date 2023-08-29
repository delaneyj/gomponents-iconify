package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Carousel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCarousel0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 11a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v24a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V11Z"/><path fill="#000" stroke="#000" d="M28 17h-8v12h8V17Z"/><path stroke="#000" stroke-linecap="round" d="M44 17h-8v12h8M4 17h8v12H4"/><path stroke="#fff" stroke-linecap="round" d="M4 13v20m40-20v20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCarousel0)"/>`),
		g.Group(children),
	)
}