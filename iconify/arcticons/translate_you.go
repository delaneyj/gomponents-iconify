package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TranslateYou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.62 24.31l5.32-12.89m5.1 12.93l-5.1-12.93m3.4 8.6h-6.95"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M32.63 25.38h6.72m-3.67 0c0 4.34-5.29 11.51-10.59 12.61"/><path d="M27.93 32.79c2.13 2.4 5.61 4.74 8.82 5.2"/></g><rect width="24.67" height="24.67" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.64" ry="3.64"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.83 30.17v8.69c0 2 1.64 3.64 3.64 3.64h17.38c2 0 3.64-1.64 3.64-3.64V21.47c0-2-1.64-3.64-3.64-3.64h-8.69"/>`),
		g.Group(children),
	)
}