package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm191-480q0-13-9-22.5t-23-9.5H543q-13 0-22.5 9.5T511 544t9.5 22.5T543 576h96v64q0 27-18.5 45.5T575 704H447q-26 0-44.5-18.5T384 640V384q0-26 18.5-45t44.5-19h128q27 0 45.5 19t18.5 45q0 13 9.5 22.5t23 9.5t22.5-9.5t9-22.5q0-53-37.5-90.5T575 256H447q-53 0-90 37.5T320 384v256q0 53 37 90.5t90 37.5h128q53 0 90.5-37.5T703 640v-96z"/>`),
		g.Group(children),
	)
}