package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChainCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopChainCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.392 14.564a2 2 0 0 1-.587-2.767l2.181-3.353a2 2 0 1 1 3.353 2.18l-2.18 3.354a2 2 0 0 1-2.767.586Zm-2.442.202a3.999 3.999 0 0 1 .179-4.06l2.18-3.353a4 4 0 1 1 6.707 4.362l-2.18 3.353a4 4 0 0 1-6.886-.302Z"/><path d="M8.444 18.503a2 2 0 0 1-.587-2.767l2.181-3.353a2 2 0 1 1 3.353 2.18l-2.18 3.354a2 2 0 0 1-2.767.586Zm-2.442.202a3.999 3.999 0 0 1 .179-4.06l2.18-3.352a4 4 0 1 1 6.707 4.361l-2.18 3.353a4 4 0 0 1-6.886-.302Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopChainCircleFilled0)"/></g>`),
		g.Group(children),
	)
}