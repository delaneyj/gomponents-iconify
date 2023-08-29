package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyDigi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.963 20.132v-7.991l4 8l4-7.988v7.988m7.481 5.579v7.604c0 1.4-1.135 2.535-2.534 2.535h0c-.7 0-1.334-.284-1.793-.742"/><rect width="5.069" height="6.717" x="20.375" y="25.72" rx="2.535" ry="2.535" transform="rotate(-180 22.91 29.079)"/><path d="M24.002 18.14v2.7a2 2 0 0 1-2 2h0a1.993 1.993 0 0 1-1.414-.585"/><path d="M24.002 14.84v3.3a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2v-3.3m-2.198 10.88v6.717m10.233-6.717v6.717m-13.005-4.182c0-1.4-1.135-2.535-2.535-2.535h0a2.535 2.535 0 0 0-2.534 2.535v1.648c0 1.4 1.134 2.534 2.534 2.534h0c1.4 0 2.535-1.135 2.535-2.535m0 2.535V22.298"/></g>`),
		g.Group(children),
	)
}