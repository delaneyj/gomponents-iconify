package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm192-735q0-14-10-23.5t-22-9.5q-16 0-27 16L512 458L380 272q-12-16-28-16q-13 0-22.5 9.5T320 289v447q0 13 9.5 22.5T352 768t22.5-9.5T384 736V388l100 140q11 15 28 15q16 0 28-15l100-140v348q0 13 9.5 22.5T672 768t22.5-9.5T704 736V289z"/>`),
		g.Group(children),
	)
}