package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmicPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.63 14.63H42.5V7.5a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h7.13Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.56 29.108a1 1 0 0 0-1 1V42.5H40.5a2 2 0 0 0 2-2V29.108Z"/><rect width="4.752" height="6.252" x="37.748" y="18.626" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="4.752" height="6.252" x="28.654" y="18.626" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="4.752" height="6.252" x="19.561" y="18.626" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/>`),
		g.Group(children),
	)
}