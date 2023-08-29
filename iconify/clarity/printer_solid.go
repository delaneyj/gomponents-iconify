package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M29 9h-2V5H9v4H7a4 4 0 0 0-4 4v11h6v8h18v-8h6V13a4 4 0 0 0-4-4Zm-4 15v6H11V19h14Zm0-15H11V7h14Zm4 6h-2v-2h2Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}