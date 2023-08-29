package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipPrevOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.75 7a.75.75 0 0 0-1.5 0v10a.75.75 0 0 0 1.5 0V7Z"/><path fill="currentColor" fill-rule="evenodd" d="M9.393 13.252a1.584 1.584 0 0 1 0-2.505a25.76 25.76 0 0 1 7.143-3.9l.466-.166c1.023-.364 2.1.329 2.238 1.381c.34 2.59.34 5.286 0 7.876c-.138 1.052-1.215 1.745-2.238 1.381l-.466-.165a25.758 25.758 0 0 1-7.143-3.902Zm.918-1.318a.084.084 0 0 0 0 .132a24.257 24.257 0 0 0 6.727 3.674l.466.166c.1.035.232-.033.249-.163c.322-2.46.322-5.025 0-7.486a.194.194 0 0 0-.25-.163l-.465.166c-2.423.86-4.694 2.1-6.727 3.674Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}