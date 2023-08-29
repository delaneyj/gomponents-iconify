package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toolbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M41.5 17a.5.5 0 0 0-.5.5v13a.5.5 0 0 1-.5.5h-39a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 0-1 0v13c0 .827.673 1.5 1.5 1.5h39c.827 0 1.5-.673 1.5-1.5v-13a.5.5 0 0 0-.5-.5z"/><path d="M40.5 7h-39C.673 7 0 7.673 0 8.5v6a.5.5 0 0 0 .5.5H14v3.5c0 .827.673 1.5 1.5 1.5h11c.827 0 1.5-.673 1.5-1.5V15h13.5a.5.5 0 0 0 .5-.5v-6c0-.827-.673-1.5-1.5-1.5zM27 18.5a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5V15h12v3.5zM41 14H1V8.5a.5.5 0 0 1 .5-.5h39a.5.5 0 0 1 .5.5V14zM15.5 0c-.827 0-1.5.673-1.5 1.5v4a.5.5 0 0 0 1 0v-4a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 .5.5v4a.5.5 0 0 0 1 0v-4c0-.827-.673-1.5-1.5-1.5h-11z"/></g>`),
		g.Group(children),
	)
}