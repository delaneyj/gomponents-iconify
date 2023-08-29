package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Colornote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.66 12.545h38.92m-38.918 0h38.917m-39 27.315h38.933l-.921 1.843H5.5Zm14.458-14.005a2.442 2.442 0 1 0 4.78-1.006l-.513-2.435a2.478 2.478 0 0 0-2.947-1.922a2.401 2.401 0 0 0-1.832 2.928Zm-7.364 3.996l-1.518-7.214l6.297 6.208l-1.518-7.214m14.03 4.67l-1.519-7.214m-2.435.513l4.78-1.006m1.804-.38l1.519 7.214m-.76-3.607l2.345-.494m-3.104-3.113l3.607-.76m-2.088 7.973l3.607-.759m-26.564 9.254l29.321-5.701M30.924 39.86s2.957-.97 4.276-1.588s3.624-2.077 3.624-2.077s1.865-2.929 2.648-4.602s1.995-5.416 1.995-5.416M4.617 6.711h38.931V39.86H4.616Z"/>`),
		g.Group(children),
	)
}