package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Praxis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 28.215v-8.43h2.76a2.831 2.831 0 0 1 0 5.663H9.5m21.429-2.818l-4.215 5.585m4.215 0l-4.215-5.585m-9.581 2.108a2.108 2.108 0 0 1 2.108-2.108h0m-2.108 0v5.585m15.776-5.585v5.585m1.991-.471a2.37 2.37 0 0 0 1.734.471h.473A1.395 1.395 0 0 0 38.5 26.82h0a1.395 1.395 0 0 0-1.393-1.396h-.946a1.395 1.395 0 0 1-1.393-1.397h0a1.395 1.395 0 0 1 1.393-1.396h.473a2.37 2.37 0 0 1 1.733.471m-13.651 3.006a2.108 2.108 0 0 1-2.107 2.107h0a2.108 2.108 0 0 1-2.108-2.107v-1.37a2.108 2.108 0 0 1 2.108-2.108h0a2.108 2.108 0 0 1 2.107 2.108m0 3.477V22.63m9.616-1.975h-2.846"/>`),
		g.Group(children),
	)
}