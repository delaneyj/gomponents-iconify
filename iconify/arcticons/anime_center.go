package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnimeCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.62 34.75v-21.5L24 2.5L5.38 13.25v21.5L24 45.5l18.62-10.75zm-4.331-19L24 7.5m0 33l14.289-8.25M9.711 15.75v16.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.678 17.203A8.254 8.254 0 0 0 16.136 21.5m-.016 4.953a8.25 8.25 0 1 0 14.676-7.13"/><circle cx="15.75" cy="24" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="29.75" cy="18.25" r="1.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}