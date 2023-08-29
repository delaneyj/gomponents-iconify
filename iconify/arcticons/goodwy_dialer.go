package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoodwyDialer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24.076" cy="23.799" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.797 16.217a16.323 16.323 0 0 1-.597-2.968c-.106-1.003-.981-1.749-1.99-1.749h-4.678c-1.203 0-2.13 1.039-2.024 2.237c1.067 12.072 10.683 21.688 22.755 22.755c1.198.106 2.237-.818 2.237-2.02V30.3c0-1.523-.746-2.395-1.749-2.501a16.304 16.304 0 0 1-2.968-.597a3.32 3.32 0 0 0-3.298.843l-2.002 2.002a21.173 21.173 0 0 1-8.53-8.531l2.001-2.002a3.32 3.32 0 0 0 .843-3.298Z"/>`),
		g.Group(children),
	)
}