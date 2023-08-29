package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EcgWaves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 36.873h9.227l5.486-7.232l5.237 13.466l5.736-15.71l5.361 17.206l6.11-7.73H61"/>`),
		g.Group(children),
	)
}