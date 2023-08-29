package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Turff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.785 8.4v25.45c0 5.426 6.295 9.65 11.046 9.65c5.47 0 8.415-.813 11.384-3.022l-5.298-5.428c-2.085 1.163-4.62 2.236-7.342-.65V18.8h8.158v-7.15h-8.158V4.5l-9.79 3.9Z"/>`),
		g.Group(children),
	)
}