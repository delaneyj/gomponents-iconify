package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TowerOfBabel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTowerOfBabel0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M20 14.5V4h8v9.5M14 25v-9.538L34 13v10M11 35v-9l26-3v9"/><path fill="#fff" d="M40 44H8v-8l32-4v12Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTowerOfBabel0)"/>`),
		g.Group(children),
	)
}