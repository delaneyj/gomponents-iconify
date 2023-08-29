package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IlryMobiili(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.911 7.814l-5.188 12.654m-4.614 8.524l-8.597 10.63m13.17-18.969c2.325 1.243 3.196 3.805 2.02 5.943c-1.176 2.137-4.022 3.163-6.6 2.38m15.634-8.343c-2.326 1.243-3.197 3.805-2.02 5.943c1.176 2.137 4.021 3.163 6.6 2.38m-14.961 9.787c-.05-1.565.814-3.047 2.298-3.94c1.484-.892 3.388-1.075 5.062-.487m3.023-13.679l-3.243-7.834M5.5 39.698l13.858-.934m15.003-9.799l8.083 11.068m-13.092-1.708L42.5 40.184"/>`),
		g.Group(children),
	)
}