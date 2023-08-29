package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Obb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2.25 2.25 0 0 0-2 2v33a2.25 2.25 0 0 0 2 2h33a2.25 2.25 0 0 0 2-2v-33a2.25 2.25 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.7 18.566a3.632 3.632 0 0 0-4.2 3.585v3.698a3.63 3.63 0 0 0 7.259 0V22.15a3.629 3.629 0 0 0-.149-1.03"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M24.976 24a2.74 2.74 0 0 1 0 5.478h-4.52V18.522h4.52a2.74 2.74 0 0 1 0 5.478Zm0 0h-4.52m15.304 0a2.74 2.74 0 0 1 0 5.478h-4.52V18.522h4.52a2.74 2.74 0 0 1 0 5.478Zm.001 0h-4.52"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.237 20.643l2.121-2.121"/>`),
		g.Group(children),
	)
}