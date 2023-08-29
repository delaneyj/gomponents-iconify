package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BnaPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.12 24a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4Zm0 0h-3.3m7.634 4v-8l5.3 8v-8m6.929 5.324h-3.465m-.864 2.653L28.953 20l2.6 8m1.627-5.891h4m-2-2v4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}