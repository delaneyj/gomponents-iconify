package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.237 20.492l-2.25 7.016l-2.118-7.016l-2.251 7.016L8.5 20.492m31 0l-5.295 7.016m5.295 0l-5.295-7.016m-10.199 5.692a2.562 2.562 0 0 1-2.25 1.324a2.655 2.655 0 0 1-2.648-2.647V23.14a2.648 2.648 0 1 1 5.295 0v.927h-5.295m8.219 2.117a2.562 2.562 0 0 0 2.25 1.324a2.655 2.655 0 0 0 2.648-2.647V23.14a2.648 2.648 0 1 0-5.295 0v.927h5.295"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}