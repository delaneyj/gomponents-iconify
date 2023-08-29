package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warnwetter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.93 43.5l16.87-14h-9.53l3.56-6.41h-7L15.56 33h7.79Zm-4.55-20.4h19.27c10.46 0 8.71-15.74-1.78-12.25c0-8.74-15.74-8.74-15.74 1.75c-8.75-1.75-8.75 10.5-1.75 10.5Z"/>`),
		g.Group(children),
	)
}