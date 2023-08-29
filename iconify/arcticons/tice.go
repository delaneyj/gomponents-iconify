package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.177 37.016c-2.126 2.74-3.605 4.852-4.581 6c-.47.609-1.503.68-2.114 0c-1.668-1.93-9.79-12.098-9.79-12.098s3.972-4.71 5.528-6.816c-4.675-5.42-.416-11.516 4.78-11.473c5.196.043 8.863 6.358 4.368 11.568l5.94 6.721c2.319-3.983 4.13-5.313 4.13-13.058S30.69 4.5 24 4.5S9.562 10.183 9.562 17.86c0 3.82.442 5.988 1.18 7.749"/>`),
		g.Group(children),
	)
}