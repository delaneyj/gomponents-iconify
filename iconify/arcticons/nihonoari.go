package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nihonoari(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 30.081a68.027 68.027 0 0 1-9.46-.558"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.107 29.712l-.018 1.832a54.48 54.48 0 0 0 7.911.492a54.48 54.48 0 0 0 7.912-.492l.005-1.834"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 32.036v1.825h-6.764M24 30.081a68.027 68.027 0 0 0 9.46-.558M24 33.861h6.764m-11.601-8.978a5.927 5.927 0 1 1 9.5.233"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.811 31.215l-7.266 6.72m36.96.026l-7.285-6.746m-3.553-3.291l-7.644-7.08l-7.655 7.08m2.578 3.932V42.5m10.108-10.644V42.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}