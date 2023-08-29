package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wachatviewer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="21.566" cy="21.566" r="15.639" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.807 32.807l2.932 2.932m0 0a1.555 1.555 0 0 1 1.955 0l4.887 4.887a1.382 1.382 0 1 1-1.955 1.955l-4.887-4.887a1.555 1.555 0 0 1 0-1.955zM21.569 12.28a9.294 9.294 0 0 0-8.141 13.759l-1.145 4.81l4.812-1.144a9.285 9.285 0 1 0 4.474-17.426z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.856 21.755c.42-.429 1.993-2.04 2.14-2.182a.848.848 0 0 0 .157-.817a16.476 16.476 0 0 1-.3-1.705a.546.546 0 0 0-.622-.48h-2.298a.405.405 0 0 0-.367.377a9.603 9.603 0 0 0 1.246 4.728v.001l.01.018l.034.06l.002-.002a9.586 9.586 0 0 0 3.517 3.509l-.004.003l.118.065l.004.003h0a9.537 9.537 0 0 0 4.68 1.222a.406.406 0 0 0 .377-.367V23.89a.546.546 0 0 0-.479-.623a16.444 16.444 0 0 1-1.704-.3a.846.846 0 0 0-.816.157c-.141.147-1.75 1.721-2.18 2.142"/>`),
		g.Group(children),
	)
}