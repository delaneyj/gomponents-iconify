package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransmissionLte(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m10.57 30l.933-2h8.993l.933 2h2.208L17 15.778V11h-2v4.778L8.363 30zM16 18.365L17.697 22h-3.393zM13.37 24h5.26l.933 2h-7.126zM10.783 9.332a7 7 0 0 1 10.434 0l-1.49 1.334a5 5 0 0 0-7.453 0z"/><path fill="currentColor" d="M7.2 6.4a11.002 11.002 0 0 1 17.6 0l-1.6 1.2a9 9 0 0 0-14.401 0Z"/>`),
		g.Group(children),
	)
}