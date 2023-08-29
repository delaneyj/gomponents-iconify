package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntercomAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 23H4a3.003 3.003 0 0 1-3-3V4a3.003 3.003 0 0 1 3-3h16a3.003 3.003 0 0 1 3 3v16a3.003 3.003 0 0 1-3 3Z" opacity=".5"/><path fill="currentColor" d="M12 19a10.8 10.8 0 0 1-6.644-2.06a1 1 0 0 1 1.288-1.532A8.987 8.987 0 0 0 12 17a8.995 8.995 0 0 0 5.361-1.595a1 1 0 0 1 1.282 1.535A10.795 10.795 0 0 1 12 19zm-6-6a1 1 0 0 1-1-1V8a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1zm4 2a1 1 0 0 1-1-1V6a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1zm4 0a1 1 0 0 1-1-1V6a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1zm4-2a1 1 0 0 1-1-1V8a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1z"/>`),
		g.Group(children),
	)
}