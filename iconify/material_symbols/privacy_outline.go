package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivacyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.85 7.025q.05 0 .075-.013T10 7q.825 0 1.413.588T12 9v.175l-2.15-2.15ZM22 17.5l-4-4v1.675l-2-2V6H8.825l-2-2H16q.825 0 1.413.588T18 6v4.5l4-4v11ZM16 20H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4l2 2H4v12h5.3v-1.05q-1.875-.275-3.088-1.675T5 12h1.425q0 1.5 1.038 2.538T10 15.575q.825 0 1.563-.362T12.8 14.2l1.025 1.025q-.6.725-1.4 1.163t-1.725.562V18H16v-2l2 2q0 .825-.588 1.413T16 20Zm5.95 1.95l-1.4 1.4l-9.625-9.625q-.225.125-.45.2T10 14q-.825 0-1.412-.588T8 12v-1.2L.65 3.45l1.4-1.4l19.9 19.9ZM9.6 12.4Zm2.825-2.825ZM9.3 18h1.4h-1.4Z"/>`),
		g.Group(children),
	)
}