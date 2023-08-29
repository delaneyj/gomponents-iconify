package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TinybitAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Zm-30-26.897h11.126m-5.563 16.794V15.603m-2.204 16.794h4.408M10.5 17.807v-2.204m11.126 2.204v-2.204"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.301 24a4.199 4.199 0 0 1 0 8.397h-6.927V15.603H33.3a4.199 4.199 0 0 1 0 8.397Zm0 0h-6.927m-2.205-8.397h2.205m-2.205 16.794h2.205"/>`),
		g.Group(children),
	)
}