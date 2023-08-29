package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudSysdigSecure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m25 14l-2.1-1c-1.7-.8-2.9-2.6-2.9-4.5V2h10v6.5c0 1.9-1.1 3.7-2.9 4.5L25 14zM22 4v4.5c0 1.2.7 2.2 1.7 2.7l1.3.6l1.3-.6c1-.5 1.7-1.6 1.7-2.7V4h-6z"/><path fill="currentColor" d="M28 16v6H4V6h12V4H4c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h8v4H8v2h16v-2h-4v-4h8c1.1 0 2-.9 2-2v-6h-2zM18 28h-4v-4h4v4z"/>`),
		g.Group(children),
	)
}