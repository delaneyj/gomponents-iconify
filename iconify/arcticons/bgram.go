package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bgram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.553 14.289l3.96-1.891c2.34-.924 3.166-.533 4.607 1.623a18.225 18.225 0 0 0-3.992 4.398c14.454 1.28-.1 19.185-4.564 25.73l2.059-18.726s-5.321.612-5.677.957c-1.573 3.174-3.094 6.302-4.72 9.49l-9.247-9.635l10.209-4.65a18.153 18.153 0 0 0 .566-5.172a146.253 146.253 0 0 1-19.132 1.27c12.386-8 26.1-17.285 25.93-3.394Z"/><circle cx="34.87" cy="14.05" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}