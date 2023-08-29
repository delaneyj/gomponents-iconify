package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Awork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.69 4.5L6.739 14.438l10.59 6.224l6.363-3.68l6.74 3.68l10.831-6.224Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.07 20.424v15.098L23.69 43.5l13.24-7.737v-15.51l-6.5 3.507v11.575l-6.74-4.23l-6.361 4.265V23.76Z"/>`),
		g.Group(children),
	)
}