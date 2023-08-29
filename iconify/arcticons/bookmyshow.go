package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmyshow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.077 28.718v-9.441l4.72 9.44l4.72-9.44v9.44m5.339 0v-4.72m3.068-4.72l-3.068 4.72l-3.186-4.72m5.447 22.347a1.54 1.54 0 0 1-2.08.72L23.28 39.19a5.532 5.532 0 0 0-10.028-4.676L6.4 31.32a1.58 1.58 0 0 1-.764-2.1c3.455-7.403 7.228-15.504 10.377-22.249L20.006 5.5l1.398 3.893l3.88-1.432l1.398 3.893l3.85-1.42l1.388 3.862l3.88-1.431l1.398 3.893l3.88-1.432l1.434 3.998Z"/>`),
		g.Group(children),
	)
}