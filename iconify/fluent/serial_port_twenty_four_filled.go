package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SerialPortTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.231 10.386a2.75 2.75 0 0 1 2.676-3.387H19.09a2.75 2.75 0 0 1 2.675 3.388l-1.074 4.502A2.75 2.75 0 0 1 18.017 17H5.978a2.75 2.75 0 0 1-2.675-2.113L2.23 10.386ZM7 11.5A.75.75 0 1 0 7 10a.75.75 0 0 0 0 1.5Zm3.25-.75a.75.75 0 1 0-1.5 0a.75.75 0 0 0 1.5 0Zm-2 3.25a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm3.25-.75a.75.75 0 1 0-1.5 0a.75.75 0 0 0 1.5 0Zm1.75.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm3.25-.75a.75.75 0 1 0-1.5 0a.75.75 0 0 0 1.5 0ZM12 11.5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm3.25-.75a.75.75 0 1 0-1.5 0a.75.75 0 0 0 1.5 0Zm1.75.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Z"/>`),
		g.Group(children),
	)
}