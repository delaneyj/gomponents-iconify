package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.89 5a3.25 3.25 0 0 0-2.906 1.797l-3.817 7.632A5.233 5.233 0 0 1 5.25 14h21.5c.74 0 1.444.153 2.083.43l-3.817-7.633A3.25 3.25 0 0 0 22.11 5H9.891ZM2 19.25A3.25 3.25 0 0 1 5.25 16h21.5A3.25 3.25 0 0 1 30 19.25v3.5A3.25 3.25 0 0 1 26.75 26H5.25A3.25 3.25 0 0 1 2 22.75v-3.5Zm23.5 3a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Z"/>`),
		g.Group(children),
	)
}