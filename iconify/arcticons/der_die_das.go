package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DerDieDas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 39v-8h1.3a4 4 0 0 1 0 8Zm0-11v-8h1.3a4 4 0 0 1 0 8Zm0-11V9h1.3a4 4 0 0 1 0 8Zm10.7-1a1.936 1.936 0 0 1-1.7 1a2.006 2.006 0 0 1-2-2v-1.3a2 2 0 0 1 4 0v.7h-4m6-.7a2.006 2.006 0 0 1 2-2m-2 0V17"/><circle cx="23" cy="20.3" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23 22.7V28m5.7-1a1.936 1.936 0 0 1-1.7 1a2.006 2.006 0 0 1-2-2v-1.3a2 2 0 0 1 4 0v.7h-4M27 37a2 2 0 0 1-4 0v-1.3a2 2 0 0 1 4 0m0 3.3v-5.3m2.2 4.9a2.364 2.364 0 0 0 1.6.4h.4a1.3 1.3 0 0 0 0-2.6h-.9a1.3 1.3 0 0 1 0-2.6h.4c.9 0 1.3.1 1.6.4"/>`),
		g.Group(children),
	)
}