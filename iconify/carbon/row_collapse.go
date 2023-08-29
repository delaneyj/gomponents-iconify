package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RowCollapse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 20H6a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2zm0 6H6v-4h20zM17 7.828l2.586 2.586L21 9l-5-5l-5 5l1.414 1.414L15 7.828V14H4v2h24v-2H17V7.828z"/>`),
		g.Group(children),
	)
}