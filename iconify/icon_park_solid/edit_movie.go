package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditMovie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEditMovie0"><g fill="none"><path fill="#fff" d="M43 9v30h-9v-8h9V17h-9V9h9ZM5 17V9h9v8H5v14h9v8H5V17Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M43 17V9h-9m9 8v14m0-14h-9M5 17V9h9m-9 8v14m0-14h9M5 31v8h9m-9-8h9m29 0v8h-9m9-8h-9m0-22v8m0-8h-4m4 30v-8m0 8h-4M14 9v8m0-8h4m-4 30v-8m0 8h4m-4-22h4m16 0h-4m4 14h-4m-16 0h4"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M24 7v4m0 6v4m0 6v4m0 6v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEditMovie0)"/>`),
		g.Group(children),
	)
}