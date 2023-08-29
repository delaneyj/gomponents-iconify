package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M325.8 0v186.2h69.8v46.5H116.4v-46.5h69.8V0H0v186.2h69.8v93.1h162.9v46.5h-69.8V512h186.2V325.8h-69.8v-46.5h162.9v-93.1H512V0H325.8zM46.5 139.6V46.5h93.1v93.1H46.5zm256 232.8v93.1h-93.1v-93.1h93.1zm163-232.8h-93.1V46.5h93.1v93.1z"/>`),
		g.Group(children),
	)
}