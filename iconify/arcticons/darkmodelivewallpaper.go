package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Darkmodelivewallpaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 11.974L13.908 35.478l-2.033-6.315l-4.365 8.692"/><circle cx="13.419" cy="15.733" r="3.578" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.94 12.254l-.915-.916m8.789 8.79l-.916-.916m-3.478-8.399V9.518m0 12.43v-1.296m4.919-4.919h1.296m-12.431 0H8.5m8.398-3.479l.916-.916m-8.789 8.79l.916-.916m25.167-7.708l-.069.005a4.225 4.225 0 0 1-2.88 7.24a4.225 4.225 0 1 0 2.949-7.246Z"/><circle cx="31.117" cy="12.443" r=".333" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="32.005" cy="16.319" r=".333" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 11.974l10.092 23.504l2.033-6.315l4.302 8.567M24 45.406V11.974"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}