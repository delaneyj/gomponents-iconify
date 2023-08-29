package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><circle cx="19" cy="9" r=".5" stroke="currentColor"/><path d="M9 8a1 1 0 0 0 0 2V8Zm4 2a1 1 0 1 0 0-2v2Zm-4 4a1 1 0 1 0 0 2v-2Zm10 2a1 1 0 1 0 0-2v2ZM9 18a1 1 0 1 0 0 2v-2Zm8 2a1 1 0 1 0 0-2v2Zm9-10v17h2V10h-2Zm-1 18H11v2h14v-2Zm-14 0a1 1 0 0 1-1-1H8a3 3 0 0 0 3 3v-2Zm15-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM25 9a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM10 27v-2H8v2h2ZM25 7h-2v2h2V7ZM9 10h4V8H9v2Zm0 6h10v-2H9v2Zm0 4h8v-2H9v2ZM7 5h14V3H7v2Zm15 1v17h2V6h-2Zm-1 18H7v2h14v-2ZM6 23V6H4v17h2Zm1 1a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3v-2Zm15-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM21 5a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM7 3a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V3Z"/></g>`),
		g.Group(children),
	)
}