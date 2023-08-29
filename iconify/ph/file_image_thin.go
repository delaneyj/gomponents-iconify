package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileImageThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M107.33 149.78a4 4 0 0 0-6.66 0l-24.25 36.37l-13.06-20.31a4 4 0 0 0-6.72 0l-36 56A4 4 0 0 0 24 228h128a4 4 0 0 0 3.33-6.22Zm-76 70.22L60 175.4l13 20.22a4 4 0 0 0 3.33 1.83a3.93 3.93 0 0 0 3.36-1.78L104 159.21L144.53 220Zm179.5-134.83l-56-56A4 4 0 0 0 152 28H56a12 12 0 0 0-12 12v96a4 4 0 0 0 8 0V40a4 4 0 0 1 4-4h92v52a4 4 0 0 0 4 4h52v124a4 4 0 0 1-4 4h-8a4 4 0 0 0 0 8h8a12 12 0 0 0 12-12V88a4 4 0 0 0-1.17-2.83ZM156 84V41.65L198.34 84Z"/>`),
		g.Group(children),
	)
}