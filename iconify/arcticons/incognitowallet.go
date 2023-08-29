package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Incognitowallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42 39.5h0a1.5 1.5 0 0 0 1.5-1.5h0"/><path fill="none" stroke="currentColor" stroke-dasharray="2.897 4.828" stroke-linecap="round" stroke-linejoin="round" d="M43.5 33.172V12.414"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 10h0A1.5 1.5 0 0 0 42 8.5h0"/><path fill="none" stroke="currentColor" stroke-dasharray="2.919 4.865" stroke-linecap="round" stroke-linejoin="round" d="M37.135 8.5H8.432"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 8.5h0A1.5 1.5 0 0 0 4.5 10h0"/><path fill="none" stroke="currentColor" stroke-dasharray="2.897 4.828" stroke-linecap="round" stroke-linejoin="round" d="M4.5 14.828v20.758"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 38h0A1.5 1.5 0 0 0 6 39.5h0"/><path fill="none" stroke="currentColor" stroke-dasharray="2.919 4.865" stroke-linecap="round" stroke-linejoin="round" d="M10.865 39.5h28.703"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 20.4h-1"/><path fill="none" stroke="currentColor" stroke-dasharray="2.067 2.067" stroke-linecap="round" stroke-linejoin="round" d="M40.433 20.4h-4.085a2.848 2.848 0 0 0-2.848 2.848v1.504a2.848 2.848 0 0 0 2.848 2.848h5.118"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 27.6h1"/><circle cx="36.472" cy="24" r=".795" fill="currentColor"/>`),
		g.Group(children),
	)
}