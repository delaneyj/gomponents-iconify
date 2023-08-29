package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 13.5a7.48 7.48 0 0 1-5-1.92V14h-2v-3.71l-9-4.61V9h-2V5.27l-9.67 4a.52.52 0 0 0-.33.48v20.7a.54.54 0 0 0 .74.49L12 27.12V23h2v4.53l9 4.61V28h2v3.79l8.63-2.7a.54.54 0 0 0 .37-.51V12.34a7.49 7.49 0 0 1-4 1.16ZM14 21h-2v-4h2Zm0-6h-2v-4h2Zm11 11h-2v-4h2Zm0-6h-2v-4h2Z"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}