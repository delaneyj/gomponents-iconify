package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBoxTopCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 30H6a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h20a2.002 2.002 0 0 1 2 2v24a2.002 2.002 0 0 1-2 2ZM6 4v24h20V4Z"/><path fill="currentColor" d="M22 9H10V7h12zm-2 5h-8v-2h8z"/>`),
		g.Group(children),
	)
}