package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keycip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.567 24.827h2.898v2.898h-2.898zm4.984 0h2.898v2.898h-2.898zm4.984 0h2.898v2.898h-2.898zm4.985 0h2.898v2.898H32.52zm-19.938 4.984h2.898v2.898h-2.898zm4.985 0h2.898v2.898h-2.898zm4.984 0h2.898v2.898h-2.898zm4.984 0h2.898v2.898h-2.898zm4.985 0h2.898v2.898H32.52z"/><rect width="32.457" height="22.024" x="7.771" y="20.476" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.012"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.316 20.476c-.403-20.261 25.103-19.672 24.717 0M17.625 36.695h12.751v2.898H17.625zm-5.043-11.868h2.898v2.898h-2.898z"/>`),
		g.Group(children),
	)
}