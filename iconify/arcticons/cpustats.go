package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cpustats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.69 31.17h8.56v7.17h0h-8.56a1.19 1.19 0 0 1-1.19-1.2v-4.78a1.19 1.19 0 0 1 1.19-1.19Zm8.56 0H24v7.17h-9.75zm9.75 0h9.75v7.17H24zm9.75 0h9.75v6a1.19 1.19 0 0 1-1.19 1.19h-8.56h0v-7.19h0ZM15.44 24H24v7.17h0h-9.75h0v-6A1.19 1.19 0 0 1 15.44 24ZM24 24h9.75v7.17H24z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.75 24h9.75v7.17h-9.75zm-8.56-7.17h8.56V24h0H24h0v-6a1.19 1.19 0 0 1 1.19-1.17Zm8.56 0h9.75V24h-9.75z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.94 9.67h7.36a1.19 1.19 0 0 1 1.19 1.19v6h0h-9.74h0v-6a1.19 1.19 0 0 1 1.19-1.19Z"/>`),
		g.Group(children),
	)
}