package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 2a7 7 0 1 0 0 14A7 7 0 0 0 8 2zm0 12.625a5.624 5.624 0 1 1 0-11.25a5.624 5.624 0 1 1 0 11.25zm6.606-10.138a3 3 0 0 0-4.98-3.321a8.008 8.008 0 0 1 4.98 3.322zM6.374 1.166a3 3 0 0 0-4.98 3.321a8.006 8.006 0 0 1 4.98-3.322z"/><path fill="currentColor" d="M8 9V5H7v5h4V9z"/>`),
		g.Group(children),
	)
}