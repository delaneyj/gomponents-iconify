package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PdfDocScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 11.672V7.5a2 2 0 0 0-2-2H32m1 37h7.5a2 2 0 0 0 2-2v-6.084m-37 2V40.5a2 2 0 0 0 2 2h9.688m0-37H7.5a2 2 0 0 0-2 2v5.172m37.778 2.675L22.832 3.541a1.749 1.749 0 0 0-2.389.64l-15.3 26.508a1.75 1.75 0 0 0 .64 2.39h0l20.445 11.806a1.749 1.749 0 0 0 2.39-.64l15.301-26.509a1.75 1.75 0 0 0-.64-2.389Zm-9.027 3.831l-13.52-7.919m15.275 13.333l-17.233-9.95m15.249 13.403L16.79 18.094"/>`),
		g.Group(children),
	)
}