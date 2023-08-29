package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm191-640q0-53-37.5-90.5T575 256H352q-14 0-23 9.5t-9 22.5v448q0 13 9 22.5t23 9.5h223q53 0 90.5-37.5T703 640V384zM575 704H400q-7 0-11.5-4.5T384 688V336q0-6 4.5-11t11.5-5h175q27 0 45.5 19t18.5 45v256q0 27-18.5 45.5T575 704z"/>`),
		g.Group(children),
	)
}