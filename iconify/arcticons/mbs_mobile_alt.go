package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MbsMobileAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.397 24.82v6.65m0-21.798v6.65m-9.193 8.008l-6.325 2.055m20.731-6.737l-6.324 2.055m-11.237-6.096l-3.909-5.38m12.814 17.635l-3.91-5.38M23 9.88l3.909-5.38M14.096 22.136l3.91-5.38m12.529-.925l6.325 2.055M16.128 11.15l6.325 2.055M11.251 43.49v-8.422l4.216 8.432l4.216-8.42v8.42m6.425-4.216a2.108 2.108 0 1 1 0 4.216H22.63v-8.432h3.478a2.108 2.108 0 1 1 0 4.216h0Zm0 0H22.63m8.703 3.292c.517.673 1.165.923 2.067.923h1.249a2.103 2.103 0 0 0 2.103-2.103v-.01a2.103 2.103 0 0 0-2.103-2.102h-1.377a2.106 2.106 0 0 1-2.106-2.106h0a2.11 2.11 0 0 1 2.11-2.11h1.242c.902 0 1.55.25 2.067.924"/>`),
		g.Group(children),
	)
}