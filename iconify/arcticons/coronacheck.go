package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coronacheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 19.5v-14h-37v14a4.5 4.5 0 0 1 0 9v14h37v-14a4.5 4.5 0 0 1 0-9Z"/><rect width="9" height="9" x="12.5" y="12.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="9" height="9" x="12.5" y="26.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="9" height="9" x="26.5" y="12.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.5 29.5h3v3h-3zm3-3h3v3h-3zm0 6h3v3h-3zm-6-6h3v3h-3zm0 6h3v3h-3z"/>`),
		g.Group(children),
	)
}