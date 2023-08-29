package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zipme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.447 5.5h33.106A1.947 1.947 0 0 1 42.5 7.447v33.106a1.947 1.947 0 0 1-1.947 1.947H7.447A1.947 1.947 0 0 1 5.5 40.553V7.447A1.947 1.947 0 0 1 7.447 5.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.368 7.447v1.948l.974.973h1.948l.974-.973V7.447m-1.948 2.921l-.974.974v1.948l.974.974l.974-.974v-1.948Zm-1.948-2.921h3.895m-2.921 0V5.5h0m1.948 1.947V5.5m-1.948 9.737V42.5m1.948-27.263V42.5m-1.948-2.921h1.948m-1.948-2.921h1.948m-1.948-2.921h1.948m-1.948-2.921h1.948m-1.948-2.921h1.948m-1.948-2.921h1.948m-1.948-2.921h1.948m-1.948-2.921h1.948m-1.948-2.922h1.948m16.552-4.868h5.842m-2.921-2.921v5.842m-4.158 21.59a.949.949 0 1 1 .949-.949a.949.949 0 0 1-.95.949Zm5.538 0a.949.949 0 1 1 .95-.949a.949.949 0 0 1-.95.949Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.366 29.996h0a6.783 6.783 0 0 1 6.783 6.783h0a.965.965 0 0 1-.964.964H25.517a.965.965 0 0 1-.965-.964h0a6.783 6.783 0 0 1 6.783-6.783Zm-5.857-1.011l1.929 2.263m9.786-2.263l-1.93 2.263"/>`),
		g.Group(children),
	)
}