package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmBluepay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 26V15.828l-3.586 3.586L0 18l6-6l6 6l-1.414 1.414L7 15.828V26h12v2H7a2.002 2.002 0 0 1-2-2zm26-8h-3v-2h-2v2h-1a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h4v2h-6v2h3v2h2v-2h1a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-4v-2h6v-2zm-3-6V6a2.002 2.002 0 0 0-2-2H12v2h14v6M2 6h5v2H2zm0-4h8v2H2z"/>`),
		g.Group(children),
	)
}