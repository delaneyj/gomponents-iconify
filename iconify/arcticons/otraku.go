package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Otraku(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.752 40.676c.62-10.293 1.17-16.607.852-28.441M24 8.48A139.64 139.64 0 0 1 4.5 7.323"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.73 7.716l-.038 3.792A111.854 111.854 0 0 0 24 12.526a111.853 111.853 0 0 0 16.308-1.018l.012-3.797"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 12.526v3.778H10.058m24.19 24.372c-.62-10.293-1.17-16.607-.852-28.441M24 8.48a139.64 139.64 0 0 0 19.5-1.156M24 12.526h0m0 3.778h13.942"/>`),
		g.Group(children),
	)
}