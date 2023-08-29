package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 .586L22.414 21L21 22.414L19.586 21H1V5h2.586l-3-3L2 .586ZM5.586 7H3v12h14.586l-2.537-2.537a5 5 0 0 1-7.012-7.012L5.586 7Zm3.885 3.885a3 3 0 0 0 4.144 4.144L9.47 10.885ZM7.882 2h8.236l1.5 3H23v14h-2V7h-4.618l-1.5-3H9.118l-.64 1.279l-1.788-.894L7.882 2Z"/>`),
		g.Group(children),
	)
}