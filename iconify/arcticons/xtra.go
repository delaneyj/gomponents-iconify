package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xtra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3.876 18.846l17.766-.738l17.836 21.528c1.66-1.565 3.43-4.181 5.022-8.915L30.34 17.715l13.15-.706a10.784 10.784 0 0 0-1.5-2.836l-15.897-.208l-9.735-8.442a7.03 7.03 0 0 0-3.395 1.82l5.335 6.622l-12.384.391a8.066 8.066 0 0 0-2.037 4.49Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.973 18.25L3.5 20.577l16.728-.663l16.067 22.562a18.437 18.437 0 0 0 3.918-3.535m-8.188-19.906l11.9-.55a4.238 4.238 0 0 0-.436-1.476M13.618 6.775a19.134 19.134 0 0 0-1.9 1.468l4.131 5.528"/>`),
		g.Group(children),
	)
}