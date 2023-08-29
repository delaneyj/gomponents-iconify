package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arlo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.142 7.127c-4.776 5.491-14.645 15.202-14.645 20.216a19.536 19.536 0 0 0 1.194 6.845s-10.427-15.68-24.833-15.68h7.264c4.814 0 5.849-.876 7.938-.876c2.408 0 4.06 1.751 5.253 3.025m17.829-2.945c-3.9 3.423-7.72 7.047-7.72 10.788c0 4.219 7.273 12.373 7.273 12.373"/>`),
		g.Group(children),
	)
}