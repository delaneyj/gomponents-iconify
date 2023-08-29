package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm160-320H384V288q0-13-9-22.5t-22.5-9.5t-23 9.5T320 288v448q0 13 9.5 22.5T352 768h320q13 0 22.5-9.5T704 736t-9.5-22.5T672 704z"/>`),
		g.Group(children),
	)
}