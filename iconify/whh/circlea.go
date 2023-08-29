package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm192-640q0-53-37.5-90.5T576 256H448q-53 0-90.5 37.5T320 384v352q0 13 9.5 22.5T352 768t22.5-9.5T384 736V592q0-6 4.5-11t11.5-5h224q7 0 11.5 5t4.5 11v144q0 13 9.5 22.5T672 768t22.5-9.5T704 736V384zm-80 128H400q-7 0-11.5-4.5T384 496V384q0-26 19-45t45-19h128q27 0 45.5 19t18.5 45v112q0 7-4.5 11.5T624 512z"/>`),
		g.Group(children),
	)
}