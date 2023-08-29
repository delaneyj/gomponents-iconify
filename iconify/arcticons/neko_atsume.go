package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NekoAtsume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.408 42.5H40.5a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h2.105"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M20.93 40.711c0 2.155-2.634 3.901-5.885 3.901s-5.885-1.746-5.885-3.9s2.635-3.902 5.885-3.902s5.886 1.747 5.886 3.901Z"/><path d="M10.51 38.15a10.975 10.975 0 0 1-1.069-4.702c0-3.37 1.573-6.461 4.192-8.874m7.382-4.1a23.567 23.567 0 0 1 6.84-.996c6.016 0 11.358 2.189 14.718 5.574m-29.03-.642c-.505-1.539-.413-3.707 1.36-6.144c2.21.363 2.55.082 6.028 2.17"/><circle cx="18.675" cy="26.95" r="1.133"/><circle cx="36.434" cy="29.131" r="1.133"/><path d="M24.185 29.592c1.111 1.145 4.076.772 4.392.261c.728-1.22-4.888-1.654-4.392-.26Zm-3.417.203c-.223 3.525 5.743 1.967 5.213.627m.621.213c1.092 2.798 4.268 2.54 5.395.854"/></g>`),
		g.Group(children),
	)
}