package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scissors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.95 4.485a1.55 1.55 0 1 0 3.1 0a1.55 1.55 0 0 0-3.1 0Zm1.55 2.45A2.45 2.45 0 1 1 4.773 5.4l.964.644a1.993 1.993 0 0 0-.02.068l-.154.55l-.353.236l-.994-.664c-.442.433-1.048.7-1.716.7ZM.95 10.5a1.55 1.55 0 1 1 3.1 0a1.55 1.55 0 0 1-3.1 0ZM2.5 8.05a2.45 2.45 0 1 0 2.277 1.545L15 2.757l-.951.1a10 10 0 0 0-3.818 1.208l-3.075 1.71a1 1 0 0 0-.476.606l-.253.906L4.224 8.76A2.442 2.442 0 0 0 2.5 8.05Zm4.644 1.165l.012.007l3.075 1.71a10 10 0 0 0 3.818 1.208l.951.1L8.81 8.1L7.145 9.215Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}