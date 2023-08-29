package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shahid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="9.102" cy="15.07" r="4.602" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.913" cy="14.95" r="4.602" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.898" cy="14.979" r="4.602" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.337 33.873V26.43h37.491v11.22h-7.802v-3.857c-9.886.02-19.844.03-29.689.079Z"/>`),
		g.Group(children),
	)
}