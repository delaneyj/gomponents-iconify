package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Memrise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1m27.262 16.886H38.5M32.762 25.01H38.5m-5.738 5.738h3.741m-3.741-5.738v11.476M21.647 21.195h4.841m-4.841-9.681h4.841m-4.841 4.841h3.157m-3.157-4.841v9.681"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 21.185v-9.671l4.841 9.681l4.841-9.667v9.667m9.636-.01v-9.671l4.841 9.681l4.841-9.667v9.667m-29 15.291V25.01h3.757a3.854 3.854 0 0 1 0 7.708H9.5m3.757.001l3.757 3.764"/><circle cx="19.972" cy="25.369" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.972 28.883v7.603m2.871-1.257a3.21 3.21 0 0 0 2.813 1.257h1.7a2.866 2.866 0 0 0 2.862-2.869h0a2.866 2.866 0 0 0-2.862-2.869h-1.878a2.866 2.866 0 0 1-2.863-2.869h0a2.866 2.866 0 0 1 2.863-2.869h1.7a3.21 3.21 0 0 1 2.813 1.258"/>`),
		g.Group(children),
	)
}