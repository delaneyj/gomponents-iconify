package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polyline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M33.162 19.463a3.5 3.5 0 0 0-3.111 1.974L9.34 64.239a3.5 3.5 0 0 0 1.626 4.674a3.5 3.5 0 0 0 4.676-1.625l19.453-40.203l35.412 9.361L84.182 78.13a3.5 3.5 0 0 0 4.418 2.234a3.5 3.5 0 0 0 2.234-4.416l-14.268-43.49a3.5 3.5 0 0 0-2.431-2.293l-40.04-10.586a3.5 3.5 0 0 0-.933-.115z" color="currentColor"/>`),
		g.Group(children),
	)
}