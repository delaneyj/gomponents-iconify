package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nfctools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.07 9.29h0a18.28 18.28 0 0 1 0 29.42m-4-24.16a11.84 11.84 0 0 1 5 9.45a11.84 11.84 0 0 1-5 9.45m-3.68-14a5.67 5.67 0 0 1 0 9.1"/><rect width="17.53" height="30.79" x="4.5" y="8.6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.58"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 12.55h17.53M4.5 35.45h17.53"/>`),
		g.Group(children),
	)
}