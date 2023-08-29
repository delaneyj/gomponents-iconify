package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 15H5v-4c0-1 1-2 2-2h18c1 0 2 1 2 2v4h-2M7 15v10c0 1 1 2 2 2h14c1 0 2-1 2-2V15M7 15h18m-9-6v18m0-18c-1.333-2.833-4.1-7.4-6.5-5c-2.4 2.4 3 5 6.5 5Zm0 0c0-4.5 4.5-7.5 6.5-5.5C25 6 20 9 16 9Z"/>`),
		g.Group(children),
	)
}