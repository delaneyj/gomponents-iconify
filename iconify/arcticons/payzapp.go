package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Payzapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.27 20.238l1.538-8h2.619a2.153 2.153 0 0 1 2.164 2.687a3.409 3.409 0 0 1-3.197 2.687h-2.62m17.065-5.373l-3.419 4l-1.881-4m1.112 7.999l.769-4m-8.305 3.977l4.132-7.976m1.063 7.999l-1.063-7.999m.707 5.323h-3.465m-5.405 15.64l1.537-8h2.619a2.153 2.153 0 0 1 2.165 2.687a3.409 3.409 0 0 1-3.198 2.687h-2.619m6.687 2.626l1.538-8h2.619a2.153 2.153 0 0 1 2.164 2.687a3.409 3.409 0 0 1-3.197 2.687H30.78"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.161 22.834h8.809l-6.454 12.928h20.581m-16.193-2.584l4.132-7.976m1.064 8l-1.064-8m.708 5.324h-3.466m1.24-15.465l.543-2.822h-2.823m-2.355 0h-2.822l-.543 2.822m4.183 5.178h2.822l.542-2.823m-8 0l-.542 2.823h2.822m3.551-5.346h-2.691l-.518 2.691h2.692l.517-2.691z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/>`),
		g.Group(children),
	)
}