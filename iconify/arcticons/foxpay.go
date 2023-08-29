package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foxpay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.995 42.5c13.715-8 17.43-14.571 16.286-31.429c-14-7.428-18.571-7.428-32.571 0C6.567 27.643 10.567 34.5 23.995 42.5Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.853 12.786c-1.987 8.546 2.171 16.407 7.428 23.428c-2.546-8.71 2.133-14.165 8-18.285c-4.882 1.06-9.782 1.052-15.428-5.143Z"/><path d="M26.281 17.929c-4.305-.11-8.21-1.622-11.153-6.497c-.483 1.07-1.046 1.873-1.138 4.246m10.005 23.965c.065-5.108 2.444-9.7 7.43-13.714c4.848-3.16 6.326-7.5 5.713-12.572c-2.34 3.613-6.582 5.597-10.285 7.429s-10.78 10.036-2.857 18.857Z"/></g>`),
		g.Group(children),
	)
}