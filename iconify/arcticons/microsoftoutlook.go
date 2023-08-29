package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microsoftoutlook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.92 25.3v-4.77h0V6.5a1 1 0 0 0-1-1H14.55a1 1 0 0 0-1 1V14m4.64 8.19a3.68 3.68 0 0 0-3.89-3.68a3.83 3.83 0 0 0-3.49 3.89v3.41a3.69 3.69 0 0 0 3.69 3.69h0a3.69 3.69 0 0 0 3.69-3.69v-3.62"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 16v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V16a2 2 0 0 0-2-2h-16a2 2 0 0 0-2 2Zm39 26.5l-19-10.68"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.67 34v8.5H43.5V24.39l-15.91 9.17"/>`),
		g.Group(children),
	)
}