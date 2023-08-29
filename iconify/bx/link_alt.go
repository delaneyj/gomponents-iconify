package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.222 19.778a4.983 4.983 0 0 0 3.535 1.462a4.986 4.986 0 0 0 3.536-1.462l2.828-2.829l-1.414-1.414l-2.828 2.829a3.007 3.007 0 0 1-4.243 0a3.005 3.005 0 0 1 0-4.243l2.829-2.828l-1.414-1.414l-2.829 2.828a5.006 5.006 0 0 0 0 7.071zm15.556-8.485a5.008 5.008 0 0 0 0-7.071a5.006 5.006 0 0 0-7.071 0L9.879 7.051l1.414 1.414l2.828-2.829a3.007 3.007 0 0 1 4.243 0a3.005 3.005 0 0 1 0 4.243l-2.829 2.828l1.414 1.414l2.829-2.828z"/><path fill="currentColor" d="m8.464 16.95l-1.415-1.414l8.487-8.486l1.414 1.415z"/>`),
		g.Group(children),
	)
}