package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrMediaPlayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.124 17.13c1.552-.11 3.214-.222 4.876-.222s3.324.111 4.876.222c1.108.11 2.216.222 3.213.333c7.314.997 12.411 3.434 12.411 6.315c0 2.77-5.097 5.209-12.41 6.317c-.998.11-2.106.221-3.214.332c-1.552.11-3.214.222-4.876.222s-3.324-.111-4.987-.222m-3.102-.332C8.597 28.987 3.5 26.66 3.5 23.778s5.097-5.208 12.41-6.316"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.316 31.868C29.32 39.292 26.881 44.5 24 44.5s-5.319-5.208-6.316-12.632c-.111-.998-.222-2.106-.333-3.214c-.11-1.551-.221-3.103-.221-4.654c0-1.773.11-3.546.222-5.097c.11-1.108.221-2.216.332-3.214C18.792 8.486 21.12 3.5 23.89 3.5s5.208 4.987 6.206 12.19m.554 3.213c.111 1.662.221 3.324.221 5.097c0 1.662-.11 3.213-.222 4.654"/>`),
		g.Group(children),
	)
}