package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cputhree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.794 320v704h-832l-192-192V320q26 0 45-19t19-45.5t-19-45t-45-18.5V0h1024v192q-27 0-45.5 18.5t-18.5 45t18.5 45.5t45.5 19zm-128-96q0-40-28-68t-68-28h-576q-40 0-68 28t-28 68v576q0 40 28 68t68 28h576q40 0 68-28t28-68V224zm-96 608h-576q-13 0-22.5-9.5t-9.5-22.5V224q0-13 9.5-22.5t22.5-9.5h576q13 0 22.5 9.5t9.5 22.5v576q0 13-9.5 22.5t-22.5 9.5z"/>`),
		g.Group(children),
	)
}