package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameOfLife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.106 4.5h3.331v3.331h-3.331zm22.52 7.507h3.331v3.331h-3.331zm-22.52 28.162h3.331V43.5h-3.331zm22.52-7.507h3.331v3.331h-3.331zM14.106 7.831H8.042v9.817h3.332v-6.485h2.732V7.83m-2.732 9.818h9.817v3.331h-9.817v-3.331m19.189 3.753v-3.753h-4.775v-2.732h-6.485v-3.331h9.816v3.753h7.507v6.064h-6.063M11.374 27.02h9.817v3.332h-9.817Zm19.189-.421v2.731h-4.775v4.775h-6.485v2.31h9.816v-3.753h7.507v-6.063h-6.063M8.042 30.352v9.817h6.064v-3.332h-2.732v-6.485H8.042"/>`),
		g.Group(children),
	)
}