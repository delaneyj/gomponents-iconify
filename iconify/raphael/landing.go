package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.322 19.49s1.903.343.3-1.868c-1.354-1.866-5.262-3.104-5.262-3.104l-4.213-8.23l-2.47-.393l.973 5.45l-3.41-1.235l-.468-2.837l-1.764-.97s-.496 2.74-.15 5.27c0 0 7.107 6.426 16.464 7.918zm-20.07 3.616v1.998h24.497v-1.998H3.25zM14 17.94a.75.75 0 1 0 1.5 0a.75.75 0 0 0-1.5 0z"/>`),
		g.Group(children),
	)
}