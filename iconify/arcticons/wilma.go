package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wilma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.982 4.5a17.794 17.794 0 0 1-5.342 3.894c-3.65 1.921-9.032 2.588-13.473 5.451a15.015 15.015 0 0 0-6.729 9.47c-1.272 4.967-.867 9.932 3.489 14.827a16.464 16.464 0 0 0 25.793-1.869c2.634-3.602 3.348-11.31 1.246-15.575a15.304 15.304 0 0 0-5.732-6.48"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.982 4.5a7.268 7.268 0 0 1-.748 9.719a10.447 10.447 0 0 1-3.83 2.168c-1.397.483-2.846.795-4.27 1.196a22.607 22.607 0 0 0-5.954 2.527a13.188 13.188 0 0 0-4.511 4.575a11.446 11.446 0 0 0-1.402 6.472a12.635 12.635 0 0 0 2.274 6.238a13.26 13.26 0 0 0 13.068 5.451"/><circle cx="24.23" cy="27.514" r="8.433" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}