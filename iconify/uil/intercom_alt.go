package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntercomAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.991 15a1 1 0 0 0 1-1V6a1 1 0 1 0-2 0v8a1 1 0 0 0 1 1Zm4 0a1 1 0 0 0 1-1V6a1 1 0 1 0-2 0v8a1 1 0 0 0 1 1Zm-8-2a1 1 0 0 0 1-1V8a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1Zm14-12h-16a3.003 3.003 0 0 0-3 3v16a3.003 3.003 0 0 0 3 3h16a3.003 3.003 0 0 0 3-3V4a3.003 3.003 0 0 0-3-3Zm1 19a1 1 0 0 1-1 1h-16a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-3.643-4.591A8.82 8.82 0 0 1 11.99 17a8.987 8.987 0 0 1-5.356-1.591a1 1 0 1 0-1.287 1.53A10.8 10.8 0 0 0 11.99 19a10.8 10.8 0 0 0 6.644-2.06a1 1 0 0 0-1.287-1.531ZM17.99 7a1 1 0 0 0-1 1v4a1 1 0 1 0 2 0V8a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}