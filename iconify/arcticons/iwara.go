package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iwara(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.021 20.748c-6.04-9.805.827-15.383 6.973-15.247c6.369.142 12.599 6.716 3.695 14.244"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.33 11.488a16.745 16.745 0 0 1 4.957-.497c1.904-.18 1.448 2.41-.54 5.15a13.262 13.262 0 0 1-5.488 4.437m-3.57-.833c1.333.82 1.998 1.191 3.57.833c11.192 8.9 3.758 20.604-2.98 21.923"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.28 42.5l-15.667-.2C.965 39.036 5.84 20.96 7.709 20.98c10.778 4.46 15.445-.72 13.313-.232"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.86 25.791c1.845 1.028 3.146 2.136 3.43 3.292c.903 3.656-1.414 6.564-5.878 7.57a14.05 14.05 0 0 1-12.683-3.71m13.19-23.928c-1.59-.01-4.48 3.284-3.883 5.347"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.755 28.064l.402 2.272l-3.348.402m5.891-2.54l-2.945-.134m3.528-11.589v.004"/><circle cx="30.304" cy="13.159" r=".75" fill="currentColor"/><circle cx="11.155" cy="31.249" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}