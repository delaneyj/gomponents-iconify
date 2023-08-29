package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="arcticonsBb0" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.48 16.98L9.261 30.643l15.592 10.551L41.4 30.012m-4.369-2.956l4.369 2.956m-11.603 1.933l7.234-4.889M5.5 40.372l4.85-3.278m19.447-5.149l4.369 2.956"/></defs><use href="#arcticonsBb0" stroke-linecap="round" stroke-linejoin="round"/><use href="#arcticonsBb0" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.52 31.02l20.219-13.663L23.147 6.806L6.6 17.988m4.369 2.956L6.6 17.988m11.603-1.933l-7.234 4.889M42.5 7.628l-4.85 3.278m-19.447 5.149l-4.369-2.956"/>`),
		g.Group(children),
	)
}