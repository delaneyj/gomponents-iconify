package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warmedal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m879.417 425l-94 94q-31-56-81-100V0h128q27 0 45.5 19t18.5 45v320q0 24-17 41zm-431 599q-87 0-160.5-43t-116.5-116.5t-43-160.5t43-160.5t116.5-116.5t160.5-43t160.5 43t116.5 116.5t43 160.5t-43 160.5t-116.5 116.5t-160.5 43zm81-418l-81-168l-81 168l-188 23l138 127l-35 182l166-90l166 90l-35-182l138-127zm47-606h64v372q-31-17-64-29V0zm-320 0h64v343q-34 12-64 29V0zm-145 519l-94-94q-17-17-17-41V64q0-26 19-45t45-19h128v419q-50 44-81 100z"/>`),
		g.Group(children),
	)
}