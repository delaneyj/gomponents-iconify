package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUpC0"><g fill="none" stroke-width="4"><path stroke="#fff" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z" clip-rule="evenodd"/><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m33 27l-9-9l-9 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUpC0)"/>`),
		g.Group(children),
	)
}