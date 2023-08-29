package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#007A5E" d="M.5.5h100v200H.5z"/><path fill="#CE1126" d="M100.5.5h100v200h-100z"/><path fill="#FCD116" d="M200.5.5h100v200h-100zm-75.362 91.76l15.675 11.388l-5.987 18.426l15.674-11.388l15.674 11.388l-5.987-18.426l15.675-11.388H156.4l-5.9-18.427l-5.899 18.427z"/></g>`),
		g.Group(children),
	)
}