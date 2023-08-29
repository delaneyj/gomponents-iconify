package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElmSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.629 1.581L4.727 3.68h5.065l-2.099-2.1H2.63Zm7.159 0l4.63 4.631V1.828a.247.247 0 0 0-.246-.247H9.788ZM14.112 8L11.58 5.468L9.047 8l2.533 2.532L14.112 8Zm-1.485 3.58l1.792 1.791V9.788l-1.792 1.792Zm.744 2.839L8 9.047L2.629 14.42H13.37ZM1.581 13.37L6.953 8L1.58 2.629V13.37ZM8 6.953L9.792 5.16H6.208L8 6.953ZM.1 1.828C.1.874.874.1 1.828.1h12.344c.954 0 1.728.774 1.728 1.728v12.344c0 .954-.774 1.728-1.728 1.728H1.828A1.728 1.728 0 0 1 .1 14.172V1.828Z"/>`),
		g.Group(children),
	)
}