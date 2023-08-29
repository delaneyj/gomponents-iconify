package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buzkill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M25.797 21.992a5.528 5.528 0 0 1 0 11.057h-9.122V10.936h9.122a5.528 5.528 0 1 1 0 11.056Zm0 0h-9.122m4.561 11.057l1.984 4.015l1.985-4.015m-8.53-15.967c-2.806.912-9.176 5.52-11.386 9.555c-1.227 2.24-1.188 4.78 1.302 5.95c3.052 1.435 5.217-.96 6.208-3.996c1.274-3.905 1.251-9.587 3.876-11.51"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M31.325 17.082c2.806.912 9.176 5.52 11.386 9.555c1.227 2.24 1.188 4.78-1.302 5.95c-3.052 1.435-5.217-.96-6.208-3.996c-1.274-3.905-1.251-9.587-3.876-11.51"/>`),
		g.Group(children),
	)
}