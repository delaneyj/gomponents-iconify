package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1023 961H878l-39-39l83-82l94 93q11 11 7 28zM545 833l288-288l68 220l-136 136zM32 353Q0 320 0 272.5T33 192L192 33Q225 1 272.5.5T352 33l449 448l-320 320zm526 205.5q18-18.5 18-44T558 471L297 210q-18-18-43.5-18T210 210t-18 43.5t18 43.5l261 262q18 18 43.5 18t43.5-18.5z"/>`),
		g.Group(children),
	)
}