package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tricount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.06 9.337l.062 4.152m0 8.723l.001 4.492c0 3.11.042 5.074.181 6.462a7.946 7.946 0 0 0 .782 3.002c.854 1.693 2.532 2.71 4.891 3.005h.003l.009.001h.01a17.574 17.574 0 0 0 6.69-.647l.127-.041m8.07-10.847l.057 3.037a17.326 17.326 0 0 0 .274 3.8a5.82 5.82 0 0 0 1.85 3.276a7.083 7.083 0 0 0 3.751 1.439h.002a18.68 18.68 0 0 0 6.548-.631l.132-.038m-39-21.011h17.689m3.237 0h17.551m-8.775-8.776v17.551"/>`),
		g.Group(children),
	)
}