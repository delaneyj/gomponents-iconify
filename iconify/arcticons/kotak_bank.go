package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KotakBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="20.5" ry="17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 36.065v-24.13M24 24c-4.093 4.925-5.447 6.551-9.7 6.551c-3.264 0-5.945-2.904-5.945-6.551c0-4.013 1.682-6.794 5.888-6.794c2.748 0 3.73.196 6.308 2.916M24 24c4.093-4.925 6.056-6.794 9.7-6.794c3.264 0 5.945 3.147 5.945 6.794c0 4.013-1.682 6.916-5.888 6.916c-2.664 0-4.2-.506-6.645-3.252"/>`),
		g.Group(children),
	)
}