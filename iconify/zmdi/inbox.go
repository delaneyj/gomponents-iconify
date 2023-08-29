package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H42q-17 0-29.5-12.5T0 341V43q0-18 12.5-30.5T42 0h299zm0 256V43H42v213h86q0 27 18.5 45.5T192 320t45.5-18.5T256 256h85zm-64-107l-85 86l-85-86h42V85h86v64h42z"/>`),
		g.Group(children),
	)
}