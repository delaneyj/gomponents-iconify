package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DbnavigatorNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m43.5 17.194l-10.873 1.863c-5.042.864-6.964 2.091-6.964 2.091l-1.654-2.363c1.926-2.88 5.706-5.643 9.308-5.146L43.5 15.046v-3.559S23.03 4.558 17.496 7.303C9.4 11.319 6.533 16.22 4.616 26.252c-.96 5.026 3.307 10.975 38.884-3.443v-5.615Z"/><path d="M4.509 27.12c.268 2.75 8.353 2.563 10.974 1.605c3.549-1.296 9.017-5.842 17.144-7.49L43.5 19.027M9.08 14.118c-1.132 3.372.528 3.442 3.986 3.49c8.609.118 15.032-8.423 8.726-8.387l-7.677.044M6.92 26.148l1.395-3.75h.61a1.263 1.263 0 0 1 1.177 1.875h0a3.007 3.007 0 0 1-2.572 1.875h-.61Z"/><path d="M13.204 24.235a.594.594 0 0 1 .554.882c-.202.5-.672.843-1.21.882h-1.455l1.311-3.527h1.455a.594.594 0 0 1 .554.881c-.202.501-.67.843-1.21.882Zm-.001 0h-1.23"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.321 41.34v-5.3m4 5.3v-3.3a2 2 0 1 0-4 0m9.739 2.291a2 2 0 0 1-1.739 1.01h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v.65h-4m9.339-2.651l-4 5.3m4 0l-4-5.3m5.632.038h2.132m-1.065-1.688v5.955a.949.949 0 0 0 1.015.993h.305"/>`),
		g.Group(children),
	)
}