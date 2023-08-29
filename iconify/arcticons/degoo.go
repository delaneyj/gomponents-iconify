package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Degoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.414 15.19a7.465 7.465 0 0 0-9.148-5.269a7.35 7.35 0 0 0-5.27 9.148c.744 2.333 2.716 3.43 3.927 5.42"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.895 35.687a9.051 9.051 0 0 1-12.8 0a9.19 9.19 0 0 1 0-12.802a9.044 9.044 0 0 1 8.388-2.43a8.952 8.952 0 0 1 5.44 4.034l6.154 9.595a8.952 8.952 0 0 0 5.44 4.033a9.045 9.045 0 0 0 8.387-2.43a9.19 9.19 0 0 0 0-12.801a9.051 9.051 0 0 0-12.8 0"/>`),
		g.Group(children),
	)
}