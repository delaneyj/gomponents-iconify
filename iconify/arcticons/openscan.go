package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openscan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.777 28.32v-8.64l5.723 8.64v-8.64M9.671 27.373a2.417 2.417 0 0 0 2.118.947h1.28a2.157 2.157 0 0 0 2.154-2.16h0A2.158 2.158 0 0 0 13.068 24h-1.413A2.158 2.158 0 0 1 9.5 21.84h0a2.157 2.157 0 0 1 2.155-2.16h1.28a2.417 2.417 0 0 1 2.117.947m8.059 4.795v.036a2.862 2.862 0 0 1-2.861 2.862h0a2.862 2.862 0 0 1-2.862-2.862v-2.916a2.862 2.862 0 0 1 2.862-2.862h0a2.862 2.862 0 0 1 2.862 2.862v.036m1.807 5.716l2.867-8.614m2.748 8.64l-2.748-8.64m1.829 5.75h-3.743"/>`),
		g.Group(children),
	)
}