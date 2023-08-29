package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m221.9 61.38l-67.74 72.31a8 8 0 0 0-2.16 5.47v55.49a8 8 0 0 1-3.56 6.66l-32 21.33A8 8 0 0 1 104 216v-76.84a8 8 0 0 0-2.16-5.47L34.1 61.38A8 8 0 0 1 40 48h176a8 8 0 0 1 5.9 13.38Z" opacity=".2"/><path d="M230.6 49.53A15.81 15.81 0 0 0 216 40H40a16 16 0 0 0-11.81 26.76l.08.09L96 139.17V216a16 16 0 0 0 24.87 13.32l32-21.34a16 16 0 0 0 7.13-13.32v-55.49l67.74-72.32l.08-.09a15.8 15.8 0 0 0 2.78-17.23ZM40 56Zm108.34 72.28a15.92 15.92 0 0 0-4.34 10.89v55.49L112 216v-76.83a15.92 15.92 0 0 0-4.32-10.94L40 56h176Z"/></g>`),
		g.Group(children),
	)
}