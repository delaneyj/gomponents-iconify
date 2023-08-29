package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlethree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm192-448q0-57-44-96q44-38 44-96q0-53-37.5-90.5T576 256H448q-53 0-90.5 37.5T320 384q0 13 9.5 22.5T352 416t22.5-9.5T384 384q0-26 18.5-45t45.5-19h128q26 0 45 19t19 45.5t-19 45t-45 18.5h-32q-13 0-22.5 9.5T512 480t9.5 22.5T544 512h32q26 0 45 19t19 45v64q0 27-19 45.5T576 704H448q-27 0-45.5-18.5T384 640q0-13-9.5-22.5T352 608t-22.5 9.5T320 640q0 53 37.5 90.5T448 768h128q53 0 90.5-37.5T704 640v-64z"/>`),
		g.Group(children),
	)
}