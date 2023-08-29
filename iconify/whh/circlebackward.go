package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlebackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm-64-756q-18-20-31-7L200 493q-8 8-8 19t8 19l217 232q13 13 31-7V268zm320 0q-18-20-31-7L520 493q-8 8-8 19t8 19l217 232q13 13 31-7V268z"/>`),
		g.Group(children),
	)
}