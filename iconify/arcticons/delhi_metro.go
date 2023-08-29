package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DelhiMetro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.094 11.799C13.484 7.937 18.458 5.5 24 5.5s10.514 2.437 13.905 6.297m0 24.405C34.515 40.063 29.542 42.5 24 42.5s-10.515-2.437-13.905-6.298"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 14a44.852 44.852 0 0 1-16.882-3.286C4.234 14.373 2.5 18.98 2.5 24s1.735 9.627 4.618 13.286C12.333 35.174 18.028 34 24 34s11.668 1.174 16.882 3.286C43.766 33.627 45.5 29.02 45.5 24s-1.735-9.627-4.618-13.286A44.853 44.853 0 0 1 24 14Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.165 28.144v-8.297l4.154 8.306l4.153-8.294v8.294m2.077 0v-8.306h2.72c1.537 0 2.783 1.249 2.783 2.79s-1.246 2.79-2.784 2.79h-2.72m2.72-.001l2.719 2.725m7.428-2.783v.034a2.752 2.752 0 0 1-2.752 2.751h0a2.752 2.752 0 0 1-2.751-2.751v-2.804a2.752 2.752 0 0 1 2.751-2.751h0a2.752 2.752 0 0 1 2.752 2.751v.034m-30.83 5.521v-8.306h1.87a3.634 3.634 0 0 1 3.633 3.634v1.038a3.634 3.634 0 0 1-3.634 3.634H8.585Zm25.466 6.981C31.391 37.537 27.867 39 24 39s-7.391-1.463-10.051-3.866m-.001-22.268C16.608 10.463 20.133 9 24 9s7.392 1.463 10.052 3.866"/>`),
		g.Group(children),
	)
}