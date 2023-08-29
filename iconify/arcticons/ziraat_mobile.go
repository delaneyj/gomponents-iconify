package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZiraatMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 20.473L12.832 11.24V5.681c0-.997 1.161-1.544 1.93-.908l7.955 6.578A3.535 3.535 0 0 1 24 14.075v6.398Zm-1.93 22.754l-7.955-6.578a3.535 3.535 0 0 1-1.283-2.724v-6.398L24 36.76v5.558c0 .997-1.161 1.544-1.93.908ZM24 32.144L12.832 22.91v-7.006L24 25.138v7.006zm-3.199-9.65L24 20.473m-11.168 7.054l3.166-2M24 20.473l11.168-9.234V5.681c0-.997-1.161-1.544-1.93-.908l-7.955 6.578A3.535 3.535 0 0 0 24 14.075v6.398Zm1.93 22.754l7.955-6.578a3.535 3.535 0 0 0 1.283-2.724v-6.398L24 36.76v5.558c0 .997 1.161 1.544 1.93.908Zm9.238-27.323L24 25.138v7.006l11.168-9.234m-8.197-4.893v4.63m0 7.04v4.617"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.964 20.793l.36-.299a2.325 2.325 0 0 0 .844-1.791v-7.464m-1.204 9.554l.215.015c.5-.133.99.243.99.76v5.959"/>`),
		g.Group(children),
	)
}