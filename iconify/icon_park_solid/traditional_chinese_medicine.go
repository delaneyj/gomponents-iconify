package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TraditionalChineseMedicine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTraditionalChineseMedicine0"><g fill="none" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-linejoin="round" d="M24 43h5.955c.083-2.737.484-4.242 1.204-4.515C38.669 35.635 44 28.434 44 20H4c0 8.251 5.103 15.323 12.357 18.294c.758.31 1.325 1.88 1.699 4.706H24Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M14.443 27.023c.36 1.187.836 2.168 1.427 2.942a11.253 11.253 0 0 0 2.14 2.104"/><path stroke="#fff" stroke-linecap="round" d="M25.214 20.026a3301.104 3301.104 0 0 0 6.622-13.708c1.05-2.21 3.18-2.742 5.145-1.902c1.965.84 2.76 3.242 1.935 4.917c-.51 1.034-2.24 4.591-5.19 10.671"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTraditionalChineseMedicine0)"/>`),
		g.Group(children),
	)
}