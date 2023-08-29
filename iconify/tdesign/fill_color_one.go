package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FillColorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 .052L23.948 12L22 13.948V14h-.052l-7.12 7.12a4 4 0 0 1-5.656 0L2.88 14.829a4 4 0 0 1 0-5.657l4.438-4.439L4.586 2L6 .586l2.733 2.733L12 .052ZM8.733 6.148l-4.438 4.438c-.39.39-.586.902-.586 1.414h17.41L12 2.88l-1.852 1.853L13.414 8L12 9.414L8.733 6.148ZM4.88 14l5.706 5.706a2 2 0 0 0 2.828 0L19.12 14H4.88Zm16.87 4.389l1.314 1.66c.581.734.581 1.848 0 2.581a1.677 1.677 0 0 1-1.314.657c-.53 0-1-.26-1.314-.657c-.581-.733-.581-1.847 0-2.581l1.314-1.66Z"/>`),
		g.Group(children),
	)
}