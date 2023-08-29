package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlej(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm160-768q-13 0-22.5 9.5T640 288v352q0 27-19 45.5T576 704H448q-27 0-45.5-18.5T384 640q0-13-9.5-22.5T352 608t-22.5 9.5T320 640q0 53 37.5 90.5T448 768h128q53 0 90.5-37.5T704 640V288q0-13-9.5-22.5T672 256z"/>`),
		g.Group(children),
	)
}