package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Square(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.444 8.333H9.556c-.675 0-1.223.548-1.223 1.223v4.888c0 .675.548 1.223 1.223 1.223h4.888c.675 0 1.223-.548 1.223-1.223V9.556c0-.675-.548-1.223-1.223-1.223M18.111 22H5.89A3.89 3.89 0 0 1 2 18.111V5.89A3.89 3.89 0 0 1 5.889 2H18.11A3.89 3.89 0 0 1 22 5.889V18.11A3.89 3.89 0 0 1 18.111 22Zm0 2A5.89 5.89 0 0 0 24 18.111V5.89A5.89 5.89 0 0 0 18.111 0H5.89A5.89 5.89 0 0 0 0 5.889V18.11A5.89 5.89 0 0 0 5.889 24H18.11Z"/>`),
		g.Group(children),
	)
}