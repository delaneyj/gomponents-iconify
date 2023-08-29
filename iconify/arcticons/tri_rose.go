package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriRose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.084 19.14A17.243 17.243 0 0 0 24 26.057a17.278 17.278 0 0 0 20.5-6.66a17.243 17.243 0 0 0-11.416-.259Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.388 32.56A17.243 17.243 0 0 0 24 26.058a17.278 17.278 0 0 0 12.67 17.439a17.244 17.244 0 0 0-3.282-10.937Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.718 36.995A17.243 17.243 0 0 0 24 26.058a17.278 17.278 0 0 0-12.67 17.439a17.244 17.244 0 0 0 9.388-6.502Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.584 26.317A17.243 17.243 0 0 0 24 26.058a17.278 17.278 0 0 0-20.5-6.66a17.243 17.243 0 0 0 9.084 6.919Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.227 15.28A17.243 17.243 0 0 0 24 26.059a17.278 17.278 0 0 0 0-21.555a17.244 17.244 0 0 0-3.773 10.778Z"/>`),
		g.Group(children),
	)
}