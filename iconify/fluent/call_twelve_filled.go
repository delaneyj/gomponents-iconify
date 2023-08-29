package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m5.966 4.49l-.827.742a4.914 4.914 0 0 0 .455 1.073a4.736 4.736 0 0 0 .722.922l1.071-.33c.6-.185 1.255.005 1.654.48l.61.726a1.47 1.47 0 0 1-.137 2.042c-.995.908-2.527 1.215-3.674.314a10.429 10.429 0 0 1-2.516-2.87A9.986 9.986 0 0 1 2.03 4.013c-.22-1.422.821-2.56 2.119-2.948c.774-.232 1.6.166 1.884.908l.335.875c.22.576.062 1.225-.402 1.641Z"/>`),
		g.Group(children),
	)
}