package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DailPad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M202.5 93.1h93.1V0h-93.1v93.1zM342.1 0v93.1h93.1V0h-93.1zM62.8 93.1h93.1V0H62.8v93.1zm139.7 139.6h93.1v-93.1h-93.1v93.1zm139.6 0h93.1v-93.1h-93.1v93.1zm-279.3 0h93.1v-93.1H62.8v93.1zm139.7 139.7h93.1v-93.1h-93.1v93.1zm139.6 0h93.1v-93.1h-93.1v93.1zm-279.3 0h93.1v-93.1H62.8v93.1zM202.5 512h93.1v-93.1h-93.1V512z"/>`),
		g.Group(children),
	)
}