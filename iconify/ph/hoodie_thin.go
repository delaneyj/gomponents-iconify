package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HoodieThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M235.33 125.78L179 41.34A12 12 0 0 0 169 36H87a12 12 0 0 0-10 5.34l-56.33 84.44a4 4 0 0 0-.53 3.27l22.39 82.11A12 12 0 0 0 54.11 220H80a12 12 0 0 0 12-12v-20h72v20a12 12 0 0 0 12 12h25.89a12 12 0 0 0 11.58-8.84l22.39-82.11a4 4 0 0 0-.53-3.27ZM80 180a4 4 0 0 1-4-4V62.89l32 18.29V136a4 4 0 0 0 8 0V85.75l10 5.72a4 4 0 0 0 4 0l10-5.72V128a4 4 0 0 0 8 0V81.18l32-18.29V176a4 4 0 0 1-4 4Zm3.62-134.22A4 4 0 0 1 87 44h82.1a4 4 0 0 1 3.33 1.78l5.91 8.87L128 83.39L77.71 54.65ZM84 208a4 4 0 0 1-4 4H54.11a4 4 0 0 1-3.86-2.95L28.34 128.7L68 69.21V176a12 12 0 0 0 12 12h4Zm121.75 1.05a4 4 0 0 1-3.86 2.95H176a4 4 0 0 1-4-4v-20h4a12 12 0 0 0 12-12V69.21l39.66 59.49Z"/>`),
		g.Group(children),
	)
}