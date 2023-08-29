package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LedStripVariantOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.11 21.46L2.39 1.73L1.11 3l1.48 1.5L2 6.91l4.03 1.01l3.05 3.05l-2.83 1.14l5.3 1.33l1.39 1.39l-8.28-2.08l-.95 3.91l14.75 3.69l2.38 2.38l1.27-1.27M7.8 16.65l-1.92-.49l.47-1.95l1.93.48l-.48 1.96m3.85.96l-1.92-.48l.47-1.95l1.93.48l-.48 1.95m3.85.97l-1.92-.49l.48-1.95l.25.06l1.43 1.43l-.24.95m4.23-2.05l2.27.57l-.33 1.37l-1.94-1.94m-3.88-3.88L14.2 11l3.55.89l-1.9.76m-5.43-6.74l-.26 1.05l2.66 2.66l6.52 1.63l.95-3.91l-13-3.25l1.37 1.37l1.76.45m5.78 1.44l1.93.49l-.48 1.95l-1.93-.48l.48-1.96m-3.85-.96l1.92.48l-.47 1.95l-1.93-.48l.48-1.95Z"/>`),
		g.Group(children),
	)
}