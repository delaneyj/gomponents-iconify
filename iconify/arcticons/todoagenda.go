package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Todoagenda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.55 5.5H7.45a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2.05-1.95ZM8.1 13.44h4v4h-4Zm4 18h-4v-4h4Zm-4 3.44h4v4h-4Zm0-25.21h31.69m-24.25 5.77h24.25M15.54 29.49h24.25m-24.25 7.42h24.25m0-12.91H8.1"/>`),
		g.Group(children),
	)
}