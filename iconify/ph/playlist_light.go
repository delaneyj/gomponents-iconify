package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M34 64a6 6 0 0 1 6-6h176a6 6 0 0 1 0 12H40a6 6 0 0 1-6-6Zm6 70h120a6 6 0 0 0 0-12H40a6 6 0 0 0 0 12Zm72 52H40a6 6 0 0 0 0 12h72a6 6 0 0 0 0-12Zm133.75-60.28a6 6 0 0 1-7.48 4L206 120.06V192a30 30 0 1 1-12-24v-56a6 6 0 0 1 7.72-5.75l40 12a6 6 0 0 1 4.03 7.47ZM194 192a18 18 0 1 0-18 18a18 18 0 0 0 18-18Z"/>`),
		g.Group(children),
	)
}