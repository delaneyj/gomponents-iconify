package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Netzclub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.65 34.496H9.5v-9.88a13.164 13.164 0 0 1 7.15-2.055Zm10.925 0h-7.15V20.384c1.973-1.898 5.044-2.052 7.15-2.052Zm10.925 0h-7.15V15.999c1.974-2.384 5.044-2.495 7.15-2.495Z"/>`),
		g.Group(children),
	)
}