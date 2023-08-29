package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H9c-.433 0-.854.14-1.2.4L3 21ZM5 5v12l2.134-1.6a1.99 1.99 0 0 1 1.2-.4H19V5H5Zm6 8.561L7.293 9.853L8.707 8.44L11 10.733l4.293-4.293l1.414 1.414L11 13.56v.001Z"/>`),
		g.Group(children),
	)
}