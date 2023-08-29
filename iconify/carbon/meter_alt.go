package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeterAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 20a13.854 13.854 0 0 0-2.23-7.529l-1.444 1.445A11.892 11.892 0 0 1 28 20zM28 9.414L26.586 8l-8.567 8.567A3.952 3.952 0 0 0 16 16a4 4 0 1 0 4 4a3.953 3.953 0 0 0-.567-2.02zM16 22a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2zm0-14a11.909 11.909 0 0 1 6.083 1.674l1.454-1.453A13.977 13.977 0 0 0 2 20h2A12.014 12.014 0 0 1 16 8z"/>`),
		g.Group(children),
	)
}