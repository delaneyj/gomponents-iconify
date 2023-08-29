package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProxmoxVirtualEnvironment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 24l9.97-11.336c3.59-3.507 6.159-2.195 8.53.053l-10.02 11.23l9.917 11.389c-3.215 3.134-6.051 2.715-8.582-.107L24 24Zm-9.815 11.23c-2.53 2.821-5.367 3.24-8.582.106l9.918-11.39L5.5 12.718c2.371-2.248 4.94-3.56 8.53-.053L24 24l-9.815 11.23Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.77 9.495c-2.171-1.87-3.886-3.965-7.293-.066L24 15.598l-5.477-6.169c-3.407-3.899-5.122-1.804-7.293.066L24 24L36.77 9.495Zm-25.54 29.01c2.171 1.87 3.885 3.965 7.293.066L24 32.402l5.477 6.169c3.407 3.899 5.122 1.804 7.293-.066L24 24L11.23 38.505Z"/>`),
		g.Group(children),
	)
}