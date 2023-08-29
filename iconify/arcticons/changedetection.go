package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Changedetection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.12 3.51C36.22 3 47.77 21.2 33.63 32.71C41.78 19.53 27 .78 10.56 11.34c3.81-5.19 8.2-7.66 12.56-7.83ZM19.6 8.26c4.38-.11 9.11 1.7 13 6.37c-13.29-8.09-32 6.8-21.32 23.11C-4.45 26.39 6.45 8.56 19.6 8.26Zm17.17 2.29c20.94 15.14-5.61 41.77-21.08 23.06C29 41.7 47.41 26.85 36.77 10.55Zm-22.28 5.3C6.34 29 21.3 47.4 37.73 36.84C22.48 57.62-4.36 31.2 14.49 15.85Z"/>`),
		g.Group(children),
	)
}