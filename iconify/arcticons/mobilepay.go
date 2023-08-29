package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobilepay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.924 19.179l4.332 10.45a21.149 21.149 0 0 1 9.041-2.112a19.235 19.235 0 0 1 8.772 2.112v-10.45a10.335 10.335 0 0 0-3.574-1.408l-2.22-5.577a20.392 20.392 0 0 0-8.825 1.733c-6.28 2.761-7.526 5.252-7.526 5.252Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.898 27.531l2.718 6.229a2.365 2.365 0 0 1-1.221 3.114l-12.438 5.428a2.365 2.365 0 0 1-3.113-1.222L7.129 14.24a2.365 2.365 0 0 1 1.222-3.113l12.437-5.429a2.365 2.365 0 0 1 3.114 1.222l2.98 6.828"/>`),
		g.Group(children),
	)
}