package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parentheses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.445 4.539l-.438.899A14.937 14.937 0 0 0 3.5 12c0 2.355.542 4.581 1.507 6.562l.438.899l-1.798.876l-.438-.899A16.936 16.936 0 0 1 1.5 12c0-2.665.614-5.19 1.71-7.438l.437-.899l1.798.876Zm14.907-.876l.439.899A16.937 16.937 0 0 1 22.5 12c0 2.665-.614 5.19-1.71 7.438l-.438.899l-1.797-.876l.438-.9A14.936 14.936 0 0 0 20.5 12c0-2.355-.542-4.581-1.507-6.562l-.438-.899l1.797-.876Z"/>`),
		g.Group(children),
	)
}