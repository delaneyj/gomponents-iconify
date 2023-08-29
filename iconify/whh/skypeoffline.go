package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skypeoffline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.695 768q0 106-75 181t-181 75q-103 0-177-71q-41 7-79 7q-91 0-174-35.5t-143-95.5t-95.5-143t-35.5-174q0-37 7-79q-71-74-71-177q0-106 75-181t181-75q103 0 177 71q41-7 79-7q91 0 174 35.5t143 95.5t95.5 143t35.5 174q0 38-7 79q71 74 71 177zm-128-256q0-104-51.5-192.5t-140-140t-192.5-51.5q-51 0-102 14q-27-36-67.5-57t-86.5-21q-80 0-136 56.5t-56 135.5q0 46 21 86.5t57 67.5q-14 51-14 102q0 104 51.5 192.5t140 140t192.5 51.5q51 0 102-14q27 36 67 57t87 21q79 0 135.5-56t56.5-136q0-46-21-86.5t-57-67.5q14-51 14-102z"/>`),
		g.Group(children),
	)
}