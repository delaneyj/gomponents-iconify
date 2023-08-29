package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 5H7a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2"/><path d="M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v0a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2zm2.993 11.75l2.747-2.815a1.9 1.9 0 0 0 0-2.632a1.775 1.775 0 0 0-2.56 0l-.183.188l-.183-.189a1.775 1.775 0 0 0-2.56 0a1.899 1.899 0 0 0 0 2.632l2.738 2.825z"/></g>`),
		g.Group(children),
	)
}