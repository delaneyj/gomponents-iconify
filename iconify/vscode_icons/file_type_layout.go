package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeLayout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#ff931e" d="M8.533 12.742h3.733v3.734H8.533zm5.6 0h3.733v3.734h-3.733zm5.6 0h3.733v3.734h-3.733zm-11.2 5.6h3.733v3.734H8.533zm5.6 0h3.733v3.734h-3.733zm5.6 0h3.733v3.734h-3.733z"/><path fill="#675f58" d="M27.367 4.92H4.632A2.637 2.637 0 0 0 2 7.552v16.9a2.637 2.637 0 0 0 2.632 2.628h22.735A2.637 2.637 0 0 0 30 24.448V7.552a2.638 2.638 0 0 0-2.633-2.632Zm1.2 19.528a1.207 1.207 0 0 1-1.2 1.2H4.632a1.207 1.207 0 0 1-1.2-1.2V9.24h25.133v15.208Z"/>`),
		g.Group(children),
	)
}