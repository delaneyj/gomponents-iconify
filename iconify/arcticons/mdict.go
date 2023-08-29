package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mdict(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.282 13.72c0 1.1-.9 2-2 2h0c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2h0c1.1 0 2 .9 2 2m0 3.299v-5.3m-6.201 5.301l-2.6-8l-2.026 6.002m.226-.703h3.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.207 5.5c4.98 1.724 8.393 4.31 10.344 7.766a10 10 0 0 0-8.955 9.925c0 5.523 4.477 10 10 10a10 10 0 0 0 9.76-7.943c3.64 2.062 5.978 5.6 8.144 9.28C38.284 20.757 31.083 9.19 13.207 5.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.55 13.266c2.58 5.491 4.866 8.723 10.805 11.982M17.738 30.47l-6.214 6.088"/>`),
		g.Group(children),
	)
}