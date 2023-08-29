package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Livedv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="36.992" cy="12.863" r="4.803" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.51 11.223h-8.813l9.781 8.76l-5.794 4.262l19.08 15.695h6.597s.966-1.369-2.201-3.813s-12.954-10.499-12.954-10.499l5.943-4.633ZM11.111 22.965l-6.535.05a1.986 1.986 0 0 0 .543 1.936L16.763 35.07l5.07-3.248Z"/>`),
		g.Group(children),
	)
}