package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circletwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm-64-448h128q53 0 90.5-37.5T704 448v-64q0-53-37.5-90.5T576 256H448q-53 0-90.5 37.5T320 384q0 13 9.5 22.5T352 416t22.5-9.5T384 384q0-26 18.5-45t45.5-19h128q27 0 45.5 19t18.5 45v64q0 27-18.5 45.5T576 512H448q-53 0-90.5 37.5T320 640v96q0 13 9.5 22.5T352 768h320q13 0 22.5-9.5T704 736t-9.5-22.5T672 704H400q-7 0-11.5-4.5T384 688v-48q0-26 18.5-45t45.5-19z"/>`),
		g.Group(children),
	)
}