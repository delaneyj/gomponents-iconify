package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungPay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 27.946V15.98h3.917c2.215 0 4.01 1.799 4.01 4.018s-1.795 4.018-4.01 4.018H11.5m25 .974v4.038a2.991 2.991 0 0 1-2.991 2.991h0a2.982 2.982 0 0 1-2.115-.876"/><path d="M36.5 20.054v4.936a2.991 2.991 0 0 1-2.991 2.99h0a2.991 2.991 0 0 1-2.991-2.99v-4.936m-2.992 4.936a2.991 2.991 0 0 1-2.99 2.99h0a2.991 2.991 0 0 1-2.992-2.99v-1.945a2.991 2.991 0 0 1 2.991-2.99h0a2.991 2.991 0 0 1 2.991 2.99m0 4.936v-7.927"/></g>`),
		g.Group(children),
	)
}