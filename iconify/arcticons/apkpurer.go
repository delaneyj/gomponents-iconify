package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apkpurer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.065 23.92h6.087m1.564 4.696l-4.52-13.912L9.5 28.616m12.272-9.232v13.912m11.685-9.549l5.043 4.87m-6.956-2.957l6.26-6.26m-16.031 5.732a3.49 3.49 0 0 0 3.477 3.484h0a3.49 3.49 0 0 0 3.478-3.484v-2.264a3.49 3.49 0 0 0-3.478-3.484h0a3.49 3.49 0 0 0-3.478 3.484m9.772-8.164v13.912"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/>`),
		g.Group(children),
	)
}