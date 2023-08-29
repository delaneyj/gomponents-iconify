package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningAThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M171.92 112.34a4 4 0 0 0-2.28-2.85l-53-23.78l14.43-60.79a4 4 0 0 0-6.83-3.64L29.06 124.14a4 4 0 0 0 1.3 6.37l53 23.78l-14.45 60.78a4 4 0 0 0 2 4.46a3.94 3.94 0 0 0 1.88.47a4 4 0 0 0 2.94-1.28l95.2-102.86a4 4 0 0 0 .99-3.52ZM80.21 202.1l11.68-49.18a4 4 0 0 0-2.25-4.57l-50.92-22.86l81.07-87.59l-11.68 49.18a4 4 0 0 0 2.25 4.57l50.92 22.86Zm155.37 12.11l-36-72a4 4 0 0 0-7.16 0l-36 72a4 4 0 1 0 7.16 3.58L174.47 196h43.06l10.89 21.79A4 4 0 0 0 232 220a4.12 4.12 0 0 0 1.79-.42a4 4 0 0 0 1.79-5.37ZM178.47 188L196 152.94L213.53 188Z"/>`),
		g.Group(children),
	)
}