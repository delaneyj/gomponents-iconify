package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stingle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.949 2.679s16.443 4.719 17.317 26.01"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.555 15.344s.687 17.093-19.252 24.613"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.837 40.258s-16.185 5.54-29.085-11.421"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.941 42.278S2.544 28.693 14.525 11.071"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.146 19.17s9.674-14.108 30.15-8.208"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.023 21.202v-5.008s-.44-2.273-5.865-2.273s-5.845 2.088-5.845 2.088v4.91l11.797 4.257v3.86l-6.043 6.044l-5.74-5.74v-3.48"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}