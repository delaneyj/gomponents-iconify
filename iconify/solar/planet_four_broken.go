package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanetFourBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M6 20.93A8 8 0 1 0 3.07 18"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M3 11.005S5.284 13 8.784 13c2.383 0 3.647-1.182 4.716-1.496c2.009-.59 3.5-.499 3.5-.499m-14 5s1.6-.091 3.757.499C7.905 16.818 9.262 18 11.82 18a10.85 10.85 0 0 0 5.181-1.333"/><path fill="currentColor" fill-rule="evenodd" d="M20.53 4.045a2.006 2.006 0 0 0-.306-.53c.326-.091.549-.097.606.003c.058.1-.058.29-.3.527ZM17.106 6.02c-.326.091-.548.097-.606-.003c-.058-.1.058-.29.3-.527a2.018 2.018 0 0 0 .306.53Zm0 0c.455-.126 1.113-.418 1.81-.82c.695-.402 1.277-.825 1.614-1.156a2 2 0 0 1-3.424 1.976Zm-.306-.53a2 2 0 0 1 3.424-1.977c-.455.127-1.113.419-1.809.82c-.696.403-1.278.826-1.615 1.157Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}