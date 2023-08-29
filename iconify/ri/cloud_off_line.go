package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.515 2.1l19.092 19.092l-1.415 1.415l-2.014-2.015A5.984 5.984 0 0 1 17 21H7A6 6 0 0 1 5.008 9.339a6.992 6.992 0 0 1 .353-2.563L2.1 3.514L3.515 2.1ZM7 9c0 .081.002.163.006.243l.07 1.488l-1.404.494A4.002 4.002 0 0 0 7 19h10c.186 0 .369-.013.548-.037L7.03 8.445C7.01 8.627 7 8.812 7 9Zm5-7a7 7 0 0 1 6.992 7.339a6.003 6.003 0 0 1 3.212 8.65l-1.493-1.494a3.999 3.999 0 0 0-5.207-5.206L14.01 9.796A5.983 5.983 0 0 1 17 9a5 5 0 0 0-7.876-4.09l-1.43-1.43A6.97 6.97 0 0 1 12 2Z"/>`),
		g.Group(children),
	)
}