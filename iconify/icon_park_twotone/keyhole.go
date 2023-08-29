package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyhole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTKeyhole0"><g fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill-rule="evenodd" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z" clip-rule="evenodd"/><path d="M24 15a5 5 0 0 0-3 9l-1 8h8l-1-8a5 5 0 0 0-3-9Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTKeyhole0)"/>`),
		g.Group(children),
	)
}