package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Krkortnu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.579 17.842L4.5 21.947v10.264h4.105"/><circle cx="8.605" cy="16.816" r="1.283" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.71" cy="32.211" r="4.105" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35" cy="32" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.816 32.21h14.368m8.211 0H43.5v-7.184L29.132 11.684H13.737"/>`),
		g.Group(children),
	)
}