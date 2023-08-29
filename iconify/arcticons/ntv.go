package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ntv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.812 14.803v15.766a2.445 2.445 0 0 0 2.545 2.628h.764m-6.441 0v-8.672a5.19 5.19 0 0 0-5.09-5.255a5.19 5.19 0 0 0-5.09 5.256v8.671m0-8.671v-5.255"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.812 19.023l5.507.004l5.091 13.927l5.09-13.927"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/>`),
		g.Group(children),
	)
}