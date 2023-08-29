package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sharing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M325.8 0v186.2h69.8v46.5H116.4v-46.5h69.8V0H0v186.2h69.8v93.1h162.9v46.5h-69.8V512h186.2V325.8h-69.8v-46.5h162.9v-93.1H512V0z"/>`),
		g.Group(children),
	)
}