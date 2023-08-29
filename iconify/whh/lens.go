package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lens(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M911 833L696 461q-12-44-42-77h353q17 65 17 128q0 180-113 321zM512 320q-15 0-30 2L656 21q112 33 199.5 112T986 320H512zM333 443L159 142q70-67 161-104.5T512 0q39 0 79 6L377 377q-30 29-44 66zm36 196H16Q0 577 0 512q0-180 113-321l214 370q12 43 42 78zm143 65q11 0 30-3l-174 302q-113-33-200-112T37 703h469q1 0 3 .5t3 .5zm179-124l174 302q-70 67-161 104.5T512 1024q-36 0-79-7l213-368q29-28 45-69z"/>`),
		g.Group(children),
	)
}