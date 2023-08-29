package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenVerticalSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.5 2H6.9l1.6.924A3 3 0 0 1 10 5.522v6.393a1.5 1.5 0 0 0 1-1.415V7a.5.5 0 0 1 .146-.354l1.708-1.707A.5.5 0 0 0 13 4.586V3.5A1.5 1.5 0 0 0 11.5 2ZM2.005 7.069A.506.506 0 0 1 2 7V3.5A2.5 2.5 0 0 1 4.5 1h7A2.5 2.5 0 0 1 14 3.5v1.086a1.5 1.5 0 0 1-.44 1.06L12 7.207V10.5a2.5 2.5 0 0 1-2.041 2.458c-.3 1.647-2.18 2.586-3.709 1.704l-2.745-1.585a3 3 0 0 1-1.5-2.598v-3.41ZM9 5.522A2 2 0 0 0 8 3.79L5.255 2.205a1.5 1.5 0 0 0-2.25 1.299v6.975a2 2 0 0 0 1 1.732l2.745 1.585a1.5 1.5 0 0 0 2.25-1.3V5.522Z"/>`),
		g.Group(children),
	)
}