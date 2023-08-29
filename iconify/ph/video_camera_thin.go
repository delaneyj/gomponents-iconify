package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M249.89 76.47a4 4 0 0 0-4.11.2L204 104.53V72a12 12 0 0 0-12-12H32a12 12 0 0 0-12 12v112a12 12 0 0 0 12 12h160a12 12 0 0 0 12-12v-32.53l41.78 27.86A4 4 0 0 0 252 176V80a4 4 0 0 0-2.11-3.53ZM196 184a4 4 0 0 1-4 4H32a4 4 0 0 1-4-4V72a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4Zm48-15.47l-40-26.67v-27.72l40-26.67Z"/>`),
		g.Group(children),
	)
}