package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M0 0v9h9V0H0zm11 0v2h2V0h-2zm2 2v2h-2v4h2V6h2V2h-2zm0 6v3H4v2H2v2h4v-2h2v2h5v-2h2v2h3v-2h2v-2h-5V8h-2zm7 5v2h6v-2h-2v-2h-2v2h-2zM2 13v-2H0v2h2zM17 0v9h9V0h-9zM2 2h5v5H2V2zm17 0h5v5h-5V2zM3 3v3h3V3H3zm17 0v3h3V3h-3zm-9 13v2h2v-2h-2zm2 2v2h-2v4h2v-2h4v-2h2v-2h2v2h-2v2h-2v4h2v-2h2v-2h1v2h-1v2h2v-2h1v-2h2v-4h-2v-2h-7v2h-4zm11 6v2h2v-2h-2zm-11 0v2h2v-2h-2zM0 17v9h9v-9H0zm2 2h5v5H2v-5zm1 1v3h3v-3H3z"/>`),
		g.Group(children),
	)
}