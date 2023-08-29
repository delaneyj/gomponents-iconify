package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlespoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M929 808L651 573q19-30 21-93q2-71-56-129q-71-71-168-90t-154 33q-53 58-33.5 155T351 617q55 55 129 55q59 0 93-20l236 277q-134 95-297 95q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199q0 163-95 296z"/>`),
		g.Group(children),
	)
}