package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerHighThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M153.76 28.41a4 4 0 0 0-4.22.43L78.63 84H32a12 12 0 0 0-12 12v64a12 12 0 0 0 12 12h46.63l70.91 55.16a4.07 4.07 0 0 0 2.46.84a3.92 3.92 0 0 0 1.76-.41A4 4 0 0 0 156 224V32a4 4 0 0 0-2.24-3.59ZM28 160V96a4 4 0 0 1 4-4h44v72H32a4 4 0 0 1-4-4Zm120 55.82L84 166V90l64-49.78Zm47-111.61a36 36 0 0 1 0 47.59a4 4 0 1 1-6-5.3a28 28 0 0 0 0-37a4 4 0 0 1 6-5.28ZM244 128a75.88 75.88 0 0 1-19.35 50.67a4 4 0 0 1-6-5.34a68 68 0 0 0 0-90.66a4 4 0 0 1 6-5.34A75.88 75.88 0 0 1 244 128Z"/>`),
		g.Group(children),
	)
}