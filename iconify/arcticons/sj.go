package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sj(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 31.71a9.87 9.87 0 0 1-8.6-4.48a4.67 4.67 0 0 0-4.49-2.41H9.59c-2.25 0-2.95-1.65-3.57-3h5.26c5.1 0 5.85 1.87 7 3.48A6.17 6.17 0 0 0 24 28.36M7.07 19.3c-2.25 0-2.95-1.65-3.57-3h9.08c4.52 0 8.55.8 8.12 7.48c-2.8-4.78-6.16-4.48-9.52-4.48ZM24 31.71a9.87 9.87 0 0 0 8.6-4.48a4.67 4.67 0 0 1 4.49-2.41h1.32c2.25 0 3-1.65 3.57-3h-5.26c-5.1 0-5.85 1.87-7 3.48A6.17 6.17 0 0 1 24 28.36m16.93-9.06c2.25 0 3-1.65 3.57-3h-9.08c-4.52 0-8.55.8-8.12 7.48c2.8-4.76 6.16-4.47 9.52-4.47Z"/>`),
		g.Group(children),
	)
}