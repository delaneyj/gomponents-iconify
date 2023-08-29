package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAutofitWidthDottedTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m14.768 15.712l1.513-1.461l-1.513-1.461a.75.75 0 1 1 1.042-1.08l1.886 1.821a1 1 0 0 1 0 1.44l-1.886 1.82a.75.75 0 0 1-1.042-1.079Zm-9.534 0L3.72 14.251l1.513-1.461a.75.75 0 0 0-1.042-1.08l-1.886 1.821a1 1 0 0 0 0 1.44l1.886 1.82a.75.75 0 0 0 1.042-1.079ZM8 14.252a.75.75 0 0 1-.75.75h-.5a.75.75 0 1 1 0-1.5h.5a.75.75 0 0 1 .75.75Zm1.75.75a.75.75 0 1 1 0-1.5h.5a.75.75 0 0 1 0 1.5h-.5Zm2.25-.75c0 .415.336.75.75.75h.5a.75.75 0 1 0 0-1.5h-.5a.75.75 0 0 0-.75.75ZM17 5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v4.25a.75.75 0 0 0 1.5 0V5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 .5.5v4.25a.75.75 0 1 0 1.5 0V5Z"/>`),
		g.Group(children),
	)
}