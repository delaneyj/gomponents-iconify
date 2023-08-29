package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DitzoZorg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 28.46V18.121h2.326a4.523 4.523 0 0 1 4.524 4.524v1.292a4.523 4.523 0 0 1-4.524 4.523Z"/><rect width="5.17" height="6.85" x="33.331" y="21.611" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.585"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.111 21.611h5.169l-5.169 6.849h5.169"/><circle cx="19.247" cy="18.444" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.247 21.611v6.849m4.002-8.982v8.982m-1.357-6.849h1.34"/>`),
		g.Group(children),
	)
}