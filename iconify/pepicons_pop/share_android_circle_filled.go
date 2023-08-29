package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAndroidCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopShareAndroidCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M8 16a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-4a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm9-1a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-4a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm0 14a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-4a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z" clip-rule="evenodd"/><path d="m9.236 12.777l-.763-1.291l6.791-4.013l.763 1.291l-6.791 4.013Zm6.021 5.752l.772-1.286l-6.286-3.772l-.772 1.286l6.286 3.772Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopShareAndroidCircleFilled0)"/></g>`),
		g.Group(children),
	)
}