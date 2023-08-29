package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorSquareAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2zm0-2.002zM4 6a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2zm0-2.002zM20 6a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2zm0-2.002zM20 22a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2zm0-2.002z"/><rect width="10" height="10" x="7" y="7" fill="currentColor" opacity=".5" rx="1"/><path fill="currentColor" d="M18.278 5a1.936 1.936 0 0 1 0-2H5.722a1.936 1.936 0 0 1 0 2zM20 18a1.976 1.976 0 0 1 1 .278V5.722a1.936 1.936 0 0 1-2 0v12.556A1.976 1.976 0 0 1 20 18zM4 18a1.976 1.976 0 0 1 1 .278V5.722a1.936 1.936 0 0 1-2 0v12.556A1.976 1.976 0 0 1 4 18zm14.278 1H5.722a1.936 1.936 0 0 1 0 2h12.556a1.936 1.936 0 0 1 0-2z" opacity=".25"/>`),
		g.Group(children),
	)
}