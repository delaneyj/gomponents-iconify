package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.935 19.266l-2.358 9.588l-2.359-9.588l-2.359 9.588L9.5 19.266" class="b"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.866 25.618a3.127 3.127 0 1 0 6.25 0v-3.236a3.197 3.197 0 0 0-3.184-3.236a3.1 3.1 0 0 0-3.066 3.236Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.5 19.266l-2.359 9.588l-2.359-9.588l-2.359 9.588l-2.358-9.588" class="b"/>`),
		g.Group(children),
	)
}