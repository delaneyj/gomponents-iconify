package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlefive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm192-448q0-53-37.5-90.5T576 448H400q-7 0-11.5-4.5T384 432v-96q0-6 4.5-11t11.5-5h272q13 0 22.5-9.5T704 288t-9.5-22.5T672 256H352q-13 0-22.5 9.5T320 288v160q0 27 18.5 45.5T384 512h192q27 0 45.5 19t18.5 45v64q0 27-18.5 45.5T576 704H448q-27 0-45.5-18.5T384 640q0-13-9.5-22.5T352 608t-22.5 9.5T320 640q0 53 37.5 90.5T448 768h128q53 0 90.5-37.5T704 640v-64z"/>`),
		g.Group(children),
	)
}