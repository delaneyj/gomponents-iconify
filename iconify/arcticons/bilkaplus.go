package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bilkaplus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.532 16.328a2.612 2.612 0 1 1 0 5.225H10.19V11.104h4.343a2.612 2.612 0 1 1 0 5.225Zm0 0h-4.341"/><ellipse cx="19.93" cy="11.271" fill="currentColor" rx=".756" ry=".75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.93 14.405v6.922m2.854-10.346v10.448m2.665-10.448v10.448m0-2.22l4.738-4.702m-3.29 3.265l3.816 3.657m7.097-2.5a2.632 2.632 0 0 1-5.264 0V17.23a2.632 2.632 0 0 1 5.264 0m0 4.31v-6.922M9.14 36.98V26.447h3.356a3.555 3.555 0 0 1 0 7.11H9.14m9.575-6.956v10.533m2.786-6.866v4.345a2.582 2.582 0 1 0 5.164 0v-4.345m-.001 4.345v2.633m2.936-.293a3.008 3.008 0 0 0 2.066.526h.516a1.712 1.712 0 0 0 0-3.423H31.02a1.712 1.712 0 0 1 0-3.423h.516c1.162 0 1.679.132 2.066.527m.258-3.773h5m-2.5-2.5v5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}