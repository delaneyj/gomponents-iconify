package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flowernew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 640h-82l45 46q37 36 37 86.5t-36.5 87t-87 36.5t-87.5-36l-45-46v82q0 53-37.5 90.5T512 1024t-90.5-37.5T384 896v-83l-46 47q-37 36-87.5 36t-87-36.5t-36.5-87t37-86.5l45-46h-81q-53 0-90.5-37.5T0 512t37.5-90.5T128 384h82l-46-46q-36-36-36-87t36-87t87-36t87 36l46 46v-82q0-53 37.5-90.5T512 0t90.5 37.5T640 128v81l46-46q36-36 86.5-36t87 36.5t36.5 87t-36 87.5l-47 46h83q53 0 90.5 37.5T1024 512t-37.5 90.5T896 640zM512 320q-80 0-136 56t-56 136t56.5 136t136 56T648 647.5t56-136T648 376t-136-56zm0 320q-53 0-90.5-37.5T384 512t37.5-90.5T512 384t90.5 37.5T640 512t-37.5 90.5T512 640z"/>`),
		g.Group(children),
	)
}