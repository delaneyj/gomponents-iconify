package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyhole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSKeyhole0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z" clip-rule="evenodd"/><path fill="#000" stroke="#000" d="M24 15a5 5 0 0 0-3 9l-1 8h8l-1-8a5 5 0 0 0-3-9Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSKeyhole0)"/>`),
		g.Group(children),
	)
}