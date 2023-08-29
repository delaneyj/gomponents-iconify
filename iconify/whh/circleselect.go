package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleselect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm300.5-685.5Q793 319 765 319t-47 20L447 609L305 467q-19-20-47-20t-47.5 19.5T191 514t20 47l187 187q20 21 50 20q29 1 49-20l315-315q20-19 20-47t-19.5-47.5z"/>`),
		g.Group(children),
	)
}