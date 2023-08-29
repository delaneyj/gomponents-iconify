package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.3 33.181l-5.878-8.236s3.594-2.06 5.878-1.947c0 0-3.295-2.06-3.257-3.107s2.883-2.77 6.776-3.37m1.461-8.049l-1.46 8.46s.599 5.242 9.66 5.654s7.974-6.252 7.974-6.814s.112-7.001-2.658-7.675a14.405 14.405 0 0 0-5.017-.337l-.824 2.808S22.71 8.022 20.8 8.022a14.633 14.633 0 0 0-3.519.45Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.52 37.76a16.408 16.408 0 0 0-3.959-4.878l5.467-8.237c-.787-.3-3.407-1.16-3.407-1.16s1.722-1.573 1.722-3.257c0-1.274-3.064-4.156-6.849-4.755M10.74 30.996a34.007 34.007 0 0 0-5.148 4.118M16.129 15.14a13.396 13.396 0 0 0 9.612 3.814c4.942 0 6.33-1.348 7.751-2.433M26.943 27.17c0 1.856-.603 4.065-3.598 4.065s-3.033-2.77-2.92-5.28c2.246.712 8.464 2.285 13.068.225c.71 2.471-.339 4.568-1.949 4.568s-2.883-1.16-2.844-3.558m-2.36 8.388s-6.664 2.32-7.413 2.546s-2.546-3.482-1.835-4.156s2.634-1.086 3.301-.412s.263 3.988.263 3.988"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}