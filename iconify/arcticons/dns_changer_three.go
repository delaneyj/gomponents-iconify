package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnsChangerThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 18.688v10.624h2.39c2.524 0 4.649-2.125 4.649-4.648v-1.328c0-2.523-2.125-4.648-4.648-4.648H10.5Zm9.98 10.624V18.688l7.04 10.624V18.688m3.074 9.429c.664.797 1.46 1.195 2.656 1.195h1.594c1.46 0 2.656-1.195 2.656-2.656S36.305 24 34.844 24h-1.727c-1.46 0-2.656-1.195-2.656-2.656s1.195-2.656 2.656-2.656h1.594c1.195 0 1.992.265 2.656 1.195"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}