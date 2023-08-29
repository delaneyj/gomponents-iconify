package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Isic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.786 16.289v15.526m30.428-5.265a5.182 5.182 0 1 1-10.364 0v-5.183a5.16 5.16 0 0 1 5.182-5.182a5.001 5.001 0 0 1 4.99 5.183m-26.656 8.689a4.652 4.652 0 0 0 3.903 1.756h2.342a3.903 3.903 0 0 0 0-7.806h-2.537a3.903 3.903 0 0 1 0-7.806h2.342a4.19 4.19 0 0 1 3.903 1.757m3.363-1.757v15.612"/>`),
		g.Group(children),
	)
}