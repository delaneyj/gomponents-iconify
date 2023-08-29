package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAaBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M90.86 50.89a12 12 0 0 0-21.72 0l-64 136a12 12 0 0 0 21.71 10.22L42.44 164h75.12l15.58 33.11a12 12 0 0 0 21.72-10.22ZM53.74 140L80 84.18L106.27 140ZM200 92c-13.85 0-24.77 3.86-32.45 11.48a12 12 0 1 0 16.9 17c3-3 8.26-4.52 15.55-4.52c9.34 0 17.19 5.16 19.37 12.1A47.32 47.32 0 0 0 200 124c-24.26 0-44 17.94-44 40s19.74 40 44 40a47.18 47.18 0 0 0 22-5.38a12 12 0 0 0 22-6.62v-60c0-22.06-19.74-40-44-40Zm0 88c-11 0-20-7.18-20-16s9-16 20-16s20 7.18 20 16s-9 16-20 16Z"/>`),
		g.Group(children),
	)
}