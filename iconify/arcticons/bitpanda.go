package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitpanda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.711 10.444v4.829a2.893 2.893 0 0 1-1.992 2.743l-8.528 2.606m-3.507-9.476l12.588-3.97a7.467 7.467 0 0 1 3.124-.332l-1.578-1.152a6.519 6.519 0 0 0-5.682-.889l-11.96 3.771Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.504 21.23s2.224-2.018 2.224-6.355V13.04a6.533 6.533 0 0 0-8.554-5.832l-12.49 3.938V35.6l3.507-1.106v-6.988l11.273-3.554a2.949 2.949 0 0 1 3.852 2.825v2.257a7.661 7.661 0 0 1-5.211 6.556l-16.928 5.337l3.507 2.572l15.214-4.797c3.287-1.045 6.925-3.866 6.925-9.853v-2.306a6.365 6.365 0 0 0-3.319-5.314m-5.268-.567L18.19 23.83V13.717l10.038-3.165a3.07 3.07 0 0 1 3.992 2.75v1.647a5.662 5.662 0 0 1-3.985 5.714"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.809 23.953v4.308a3.494 3.494 0 0 1-2.453 3.345l-12.672 3.995V11.146l-3.507-2.572v32.354"/>`),
		g.Group(children),
	)
}