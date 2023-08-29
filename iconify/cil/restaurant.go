package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Restaurant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M224 160h-48V48h-32v112H96V48H64v124c0 45.505 34.655 83.393 80 90.715V472h32V262.715c45.345-7.322 80-45.21 80-90.715V48h-32Zm-64 72c-27.811 0-51.524-16.722-60.33-40h120.66c-8.806 23.278-32.519 40-60.33 40ZM413.567 40.187A138.648 138.648 0 0 0 296 177.224V344h104v128h32V37.351ZM400 312h-72V177.224a105.986 105.986 0 0 1 72-100.911Z"/>`),
		g.Group(children),
	)
}