package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mporezna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.189 40.4l-6.833 3.1V21.566l6.833-3.1V40.4z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.208 14.569h0c4.045-2.03 6.33-5.252 1.937-7.72c-7.507-4.215-13.022-1.953-18.501.533l-1.744.791h0l-5.09 2.31l25.546 11.083l6.833-3.1l-8.981-3.897Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.788 12.239c1.023-.517 2.523-1.9.85-2.829c-1.8-1-3.404-.715-5.815.653l4.965 2.176Zm3.282 24.715V25.396c0-.82-.486-1.562-1.238-1.888L9.385 15.505a1.126 1.126 0 0 0-1.574 1.033v15.879l19.823 8.6a1.027 1.027 0 0 0 1.437-.942v-.63m0-2.491l-4.865-2.119"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.81 29.899v-5.971a1.398 1.398 0 0 1 1.953-1.283l2.114.915a2.706 2.706 0 0 1 1.63 2.484v6.323m.001 0v-5.971a1.398 1.398 0 0 1 1.954-1.283l2.114.915a2.706 2.706 0 0 1 1.63 2.484v6.323"/>`),
		g.Group(children),
	)
}