package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm192-384V384q0-53-37.5-90.5T576 256H448q-53 0-90.5 37.5T320 384v256q0 53 37.5 90.5T448 768h128q35 0 66-18l9 9q9 9 22 9t22-9t9-21.5t-9-22.5l-9-9q18-31 18-66zm-140-55q-9-9-21.5-9t-21.5 9t-9 21.5t9 21.5l73 73q-10 3-18 3H448q-26 0-45-18.5T384 640V384q0-26 19-45t45-19h128q26 0 45 19t19 45v256q0 8-3 18z"/>`),
		g.Group(children),
	)
}