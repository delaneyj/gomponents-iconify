package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Varobank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.71 19.23l-3.16 9.54l-3.16-9.54"/><rect width="4.77" height="6.32" x="30.84" y="22.45" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.73 24.83a2.38 2.38 0 0 1 2.38-2.38h0m-2.38 0v6.32m-2.44-2.38a2.39 2.39 0 0 1-2.39 2.38h0a2.39 2.39 0 0 1-2.38-2.38v-1.56a2.39 2.39 0 0 1 2.38-2.38h0a2.39 2.39 0 0 1 2.39 2.38m0 3.94v-6.32"/>`),
		g.Group(children),
	)
}