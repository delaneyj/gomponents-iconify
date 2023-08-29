package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lotuss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.665 23.473v-1.157m-14.598.284h2.199M20.167 21v6c0 .6.4 1 .999 1h.3m14.234-.4c.4.299.8.4 1.6.4h.4c.7 0 1.3-.6 1.3-1.3h0c0-.7-.6-1.3-1.3-1.3h-.9a1.323 1.323 0 0 1-1.299-1.3h0c0-.7.6-1.3 1.3-1.3h.4c.9 0 1.299.1 1.6.4m-10.167 4.4c.4.299.8.4 1.599.4h.4c.7 0 1.3-.6 1.3-1.3h0c0-.7-.6-1.3-1.3-1.3h-.9a1.323 1.323 0 0 1-1.3-1.3h0c0-.7.6-1.3 1.3-1.3h.4c.9 0 1.3.1 1.6.4M26.849 26v1.998M22.85 22.6v3.3c0 1.099.9 1.999 2 1.999h0c1.1 0 2-.9 2-1.999v-3.3M16.083 28h0a2.005 2.005 0 0 1-2-2v-1.3c0-1.1.9-2 2-2h0c1.1 0 2 .9 2 1.999v1.3c0 1.1-.9 2-2 2ZM9 20.001v7.998h3.999"/>`),
		g.Group(children),
	)
}