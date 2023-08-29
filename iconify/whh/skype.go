package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skype(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.695 768q0 106-75 181t-181 75q-103 0-177-71q-41 7-79 7q-91 0-174-35.5t-143-95.5t-95.5-143t-35.5-174q0-37 7-79q-71-74-71-177q0-106 75-181t181-75q103 0 177 71q41-7 79-7q91 0 174 35.5t143 95.5t95.5 143t35.5 174q0 38-7 79q71 74 71 177zm-512-448q53 0 90.5 18.5t37.5 45t19 45.5t45.5 19t45-19t18.5-45q0-80-75-136t-181-56t-181 56t-75 136t75 136t181 56q53 0 90.5 18.5t37.5 45.5t-37.5 45.5t-90.5 18.5t-90.5-19t-37.5-45.5t-18.5-45t-45-18.5t-45.5 18.5t-19 45.5q0 80 75 136t181 56t181-56t75-136t-75-136t-181-56q-53 0-90.5-19t-37.5-45.5t37.5-45t90.5-18.5z"/>`),
		g.Group(children),
	)
}