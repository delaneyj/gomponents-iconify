package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timetree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.929 4.5c10.111 3.29 15.755 8.243 14.253 18.436C12.041 22.286 9.32 14.408 9.929 4.5Zm28.375 2.874c-.466 8.82-3.227 14.413-11.728 15.398c-2.088-9.88 3.648-13.757 11.728-15.398ZM4.92 32.141c5.137-7.184 10.474-10.414 18.17-6.671c-3.555 9.453-10.476 9.635-18.17 6.671Zm38.16-.961c-7.328 1.744-12.57.832-15.425-5.886c7.582-4.09 12.136-.33 15.425 5.886ZM25.191 43.5c-4.07-6.34-4.943-11.588.453-16.504C32.01 32.8 29.97 38.34 25.19 43.5Z"/>`),
		g.Group(children),
	)
}