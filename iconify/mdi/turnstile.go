package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Turnstile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 22h-6V11l-6-6V2h12v20M9.17 6.17C8.42 6.92 8 7.94 8 9H2v2h6.55c.35.6.85 1.1 1.45 1.45V19h2v-6c1.06 0 2.08-.42 2.83-1.17L9.17 6.17Z"/>`),
		g.Group(children),
	)
}