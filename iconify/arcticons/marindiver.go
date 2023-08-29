package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marindiver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".904" d="M25.33 6.67a18.17 18.17 0 0 0-16.129 9.893a18.136 18.136 0 0 1 30.47 13.33h0a18.168 18.168 0 0 1-2.038 8.275A18.14 18.14 0 0 0 25.33 6.671Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".904" d="M14.485 25.34c-3.53-2.145-6.655 3.625-6.655 3.625s-2.293-9.586 7.07-12.094c0 0-3.577 6.797 0 6.797s6.746 4.493 7.24 5.463s2.433 5.833 5.396 5.35a4.634 4.634 0 0 1 4.884 2.338a4.107 4.107 0 0 1 2.199 1.647l-2.14.074a8.886 8.886 0 0 1-5.668 2.75c-3.845.473-7.925-3.539-9.255-4.853s-.38-1.7-5.218-1.219A10.371 10.371 0 0 1 4.5 32.793c6.138-.28 13.514-5.309 9.984-7.453Zm2.852 15.135s-5.454-.332-7.278-1.905a4.369 4.369 0 0 0 2.254-.741c.755-.67 1.06.796 5.024 2.646Z"/><circle cx="29.895" cy="37.303" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}