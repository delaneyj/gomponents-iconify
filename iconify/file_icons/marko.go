package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marko(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m108.713 196.446l-36.146-58.963L0 255.99l85.368 139.502h72.258L72.21 255.99l36.504-59.544zm134.731-58.928l-23.578 38.513l36.138 58.959l23.557-38.473l-36.117-59zM85.419 116.507l85.472 139.483h72.257l-85.507-139.483H85.418zm243.093 0h-72.206l85.469 139.483l-85.469 139.502h72.257l85.471-139.502l-85.522-139.483zm98.018 0h-72.26l85.472 139.483l-85.472 139.502h72.259L512 255.99l-85.47-139.483z"/>`),
		g.Group(children),
	)
}