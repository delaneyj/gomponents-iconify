package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fuelflash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.9 5.52a3 3 0 0 0-3 3v30.1h-.46a1.91 1.91 0 1 0 0 3.82h23.2a1.91 1.91 0 1 0 0-3.82h-.41v-17h1a1.28 1.28 0 0 1 1.29 1v11.11a4.9 4.9 0 0 0 4.88 4.88h.21a4.91 4.91 0 0 0 4.88-4.88V15.84h0v-3.77a1.93 1.93 0 0 0-.57-1.36L37.28 6.1a1.91 1.91 0 1 0-2.69 2.72l2.62 2.6l-2 2a1.7 1.7 0 0 0 1.21 2.89h2.2V22h0v11.73a1 1 0 0 1-1 1h-.21a1 1 0 0 1-1-1V23.8h0V23a5.12 5.12 0 0 0-5.09-5.09h-1V8.56a3 3 0 0 0-3-3H10.9Zm2.49 3.85H24.7a1.72 1.72 0 0 1 1.73 1.72v8.11a1.72 1.72 0 0 1-1.73 1.73H13.39a1.72 1.72 0 0 1-1.73-1.73v-8.11a1.72 1.72 0 0 1 1.73-1.72Zm8.74 14.69a.31.31 0 0 1 .26.43l-1.8 3.72l1.84-.57a.3.3 0 0 1 .36.41L21 31.86l1.51-.31a.3.3 0 0 1 .28.5l-5.08 5.59a.3.3 0 0 1-.51-.28l.78-3l-1.57.6a.3.3 0 0 1-.39-.39l1.45-3.72l-2 .79a.3.3 0 0 1-.39-.4l2.31-5.37a.33.33 0 0 1 .17-.16L22 24.08h.11Zm8.1-6.2v4.04"/>`),
		g.Group(children),
	)
}