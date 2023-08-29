package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XPlore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.68 13.015H24.77c-2-.1-5.93-4.23-8.19-4.23h-9.9a2.18 2.18 0 0 0-2.18 2.18v7.34h39v-3.42a1.83 1.83 0 0 0-1.79-1.87Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 18.295h-39v18.72a2.18 2.18 0 0 0 2.16 2.2h34.66a2.18 2.18 0 0 0 2.18-2.18v-.02Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.262 39.215l2.55-2.553c.268-.223.557-.457.875-.707c2.692-2.116 2.273 1.138 6.88-1.81a6.043 6.043 0 0 0 2.188-2.337a6.24 6.24 0 0 0 .031-5.945l-4.506 4.506l-2.467-.46l-.458-2.468l4.504-4.504a4.667 4.667 0 0 0-3.082-.685a7.082 7.082 0 0 0-2.87.72a6.052 6.052 0 0 0-2.33 2.184c-2.948 4.608.306 4.188-1.81 6.88c-.253.323-.488.613-.713.884l-6.296 6.295m3.749 0l4.523-4.522"/>`),
		g.Group(children),
	)
}