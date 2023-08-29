package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlev(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm192-736q0-13-9-22.5t-22.5-9.5t-23 9.5T640 288v277L512 693L384 565V288q0-13-9-22.5t-22.5-9.5t-23 9.5T320 288v288q0 12 9 20l155 155q14 17 28 17q11 0 28-17l155-155q9-8 9-20V288z"/>`),
		g.Group(children),
	)
}