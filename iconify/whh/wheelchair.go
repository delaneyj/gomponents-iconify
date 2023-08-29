package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wheelchair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1006 875l-63 64q-1 0-1 .5t-1 1.5l-1.5 1l-2.5 1q-4 4-7 5q-18 12-40.5 10.5T851 941L678 767H384q-26 0-45-18.5T320 704V384q0-27 19-45.5t45.5-18.5t45 18.5T448 384v256h249q7-1 16 1q4 0 7 1q7 2 9 3q12 5 22 15l145 145l20-20q19-19 45.5-19t45 19t18.5 45.5t-19 44.5zM384 256q-53 0-90-37.5T257 128t37-90.5T384 0t90.5 37.5T512 128t-37.5 90.5T384 256zm146 593l3 3q19-18 45.5-18t45 18.5t18.5 45t-18 45.5q-10 10-26 15q-74 50-161.5 62t-173-16T113 911T20 760.5t-16-173T66 426q5-16 15-26q19-19 45.5-19t45 19t18.5 45.5t-18 45.5l3 3q-53 74-45.5 167t74 159.5t159.5 74T530 849z"/>`),
		g.Group(children),
	)
}