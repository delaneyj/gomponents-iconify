package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fgts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.248 19.5h5.962m-2.981 9v-9m4.977 8.014c.552.718 1.244.986 2.207.986h1.332a2.245 2.245 0 0 0 2.246-2.245v-.01A2.245 2.245 0 0 0 35.745 24h-1.47a2.248 2.248 0 0 1-2.247-2.248h0A2.252 2.252 0 0 1 34.28 19.5h1.326c.962 0 1.654.268 2.206.986M10.009 24h2.925m-2.925 4.5v-9h4.5m7.957 2.981a2.981 2.981 0 0 0-2.981-2.98h0a2.981 2.981 0 0 0-2.982 2.98v3.038a2.981 2.981 0 0 0 2.982 2.981h0a2.981 2.981 0 0 0 2.98-2.981h-2.98"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}