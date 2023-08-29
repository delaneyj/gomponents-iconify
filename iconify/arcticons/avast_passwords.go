package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AvastPasswords(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23 27.226l2.126 2.057h2.635v2.827h2.79v2.751h4.189v-3.96l-7.52-7.974l.056-.106a7.323 7.323 0 0 0-3.017-8.594a7.153 7.153 0 0 0-8.957 1.116a7.343 7.343 0 0 0-.86 9.078a7.168 7.168 0 0 0 8.585 2.813Z"/><ellipse cx="17.913" cy="17.79" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.747" ry="1.76"/>`),
		g.Group(children),
	)
}