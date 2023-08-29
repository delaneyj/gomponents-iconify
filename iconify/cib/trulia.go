package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trulia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0c-5.079.027-8.36 3.437-8.36 8.593c0 4.267 2.781 7.923 4.371 13.469c.911 3.177 1.401 7.068 1.651 9.937h4.677c.255-2.869.74-6.76 1.651-9.937c1.589-5.547 4.371-9.199 4.371-13.469c0-5.156-3.281-8.567-8.36-8.593zm.027 13.136a3.828 3.828 0 0 1-3.824-3.824a3.827 3.827 0 0 1 7.652 0a3.829 3.829 0 0 1-3.828 3.824z"/>`),
		g.Group(children),
	)
}