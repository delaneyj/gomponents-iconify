package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSatelliteServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 26a2.97 2.97 0 0 0-1.855.66L25 25.423v-1.606a3 3 0 1 0-2 0v1.606l-2.145 1.239A2.97 2.97 0 0 0 19 26a3.02 3.02 0 1 0 2.925 2.353L24 27.154l2.075 1.198A2.998 2.998 0 1 0 29 26Zm-10 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm5-10a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm5 10a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><circle cx="7" cy="20" r="2" fill="currentColor"/><path fill="currentColor" d="M14 20a4 4 0 1 1 4-4a4.012 4.012 0 0 1-4 4Zm0-6a2 2 0 1 0 2 2a2.006 2.006 0 0 0-2-2Z"/><circle cx="21" cy="12" r="2" fill="currentColor"/><path fill="currentColor" d="M13.02 28.271L3 22.427V9.574l11-6.416l11.496 6.706l1.008-1.728l-12-7a1 1 0 0 0-1.008 0l-12 7A1 1 0 0 0 1 9v14a1 1 0 0 0 .496.864L12.013 30Z"/>`),
		g.Group(children),
	)
}