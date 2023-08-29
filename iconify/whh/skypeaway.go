package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skypeaway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.695 768q0 106-75 181t-181 75q-103 0-177-71q-41 7-79 7q-91 0-174-35.5t-143-95.5t-95.5-143t-35.5-174q0-38 7-79q-71-74-71-177q0-106 75-181t181-75q103 0 177 71q41-7 79-7q91 0 174 35.5t143 95.5t95.5 143t35.5 174q0 38-7 79q71 74 71 177zm-290-190l-158-103V320q0-26-19-45t-45.5-19t-45 19t-18.5 45v194q-1 41 34 60l185 121q24 14 51 7t41-31.5t6.5-51.5t-31.5-41z"/>`),
		g.Group(children),
	)
}