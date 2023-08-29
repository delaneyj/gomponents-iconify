package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungGiftIndonesia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 20.76V40.5a2 2 0 0 1-2 2h-29a2 2 0 0 1-2-2V20.76m-2-6.37h37v6.369h-37zm18.5 0c-5.748-8.874-9.112-10.893-12.244-7c-3.27 4.064 4.029 7 12.244 7Zm0 0l-5.201-2.649M24 14.39c5.748-8.874 9.112-10.893 12.244-7c3.27 4.064-4.029 7-12.244 7Zm0 0l5.201-2.649"/>`),
		g.Group(children),
	)
}