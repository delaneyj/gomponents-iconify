package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlevimeo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm144-768q-46 0-78.5 43.5T544 398q16-14 36-14q21 0 32.5 10.5T624 432q0 25-17.5 63.5T566 563t-38 29q-7 0-14-16t-14-47t-12-59t-12.5-70.5T464 336q-15-80-64-80q-72 0-144 128l16 32q11-13 26-22.5t22-9.5t12.5 4.5t8.5 13t5 15.5t3.5 17t2.5 14q4 18 12 66t16.5 86t22.5 79.5t37.5 65T496 768q39 0 104.5-68.5T717 538t51-154q0-128-112-128z"/>`),
		g.Group(children),
	)
}