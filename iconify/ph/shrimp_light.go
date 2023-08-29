package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShrimpLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M138 116a10 10 0 1 1 10 10a10 10 0 0 1-10-10Zm83.9-32a78.1 78.1 0 0 1-77.9 74h-32a14 14 0 0 0 0 28h56a6 6 0 0 1 0 12h-50v20h34a6 6 0 0 1 0 12H96a78 78 0 0 1 0-156h116a14 14 0 0 0 0-28h-84a22 22 0 0 1-22-22a6 6 0 0 1 12 0a10 10 0 0 0 10 10h84a26 26 0 0 1 9.9 50ZM50.65 199.88l35.48-25.34A23.74 23.74 0 0 1 86 172a25.92 25.92 0 0 1 1.46-8.57L31.51 138a65.8 65.8 0 0 0 19.14 61.88ZM90.3 186.3l-29.81 21.29A65.56 65.56 0 0 0 96 218h10v-20.71a26.05 26.05 0 0 1-15.7-10.99Zm15.7-39.59V86H96a66.1 66.1 0 0 0-60.86 40.47L94 153.24a25.86 25.86 0 0 1 12-6.53ZM209.73 86H118v60h26a66.09 66.09 0 0 0 65.73-60Z"/>`),
		g.Group(children),
	)
}