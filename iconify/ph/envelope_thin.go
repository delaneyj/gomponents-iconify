package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 52H32a4 4 0 0 0-4 4v136a12 12 0 0 0 12 12h176a12 12 0 0 0 12-12V56a4 4 0 0 0-4-4Zm-96 86.57L42.28 60h171.44ZM104.63 128L36 190.91V65.09Zm5.92 5.43L125.3 147a4 4 0 0 0 5.4 0l14.75-13.52L213.72 196H42.28Zm40.82-5.43L220 65.09v125.82Z"/>`),
		g.Group(children),
	)
}