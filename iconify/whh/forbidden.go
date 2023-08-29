package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forbidden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1023q-104 0-199-40.5t-163.5-109T40.5 710T0 511t40.5-198.5t109-163T313 40.5T512 0t199 40.5t163.5 109t109 163T1024 511t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm222-199L512 602L290 824q100 71 222 71t222-71zM128 511q0 122 70 221l222-222l-221-221q-71 100-71 222zm163-313l221 220l221-220q-100-71-221-71t-221 71zm534 91L604 510l222 222q70-99 70-221t-71-222z"/>`),
		g.Group(children),
	)
}