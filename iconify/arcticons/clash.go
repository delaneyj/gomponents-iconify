package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.19 42.5a89.044 89.044 0 0 1-14.681-1.573s1.431-28.554 5.411-35.39c-.13-.297 2.992 1.212 4.422 6.266a25.557 25.557 0 0 1 4.847-.47"/><ellipse cx="21.24" cy="20.309" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.671" ry="2.13"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.19 42.5a89.044 89.044 0 0 0 14.681-1.573s-1.431-28.554-5.413-35.39c.03-.2-3.59 1.755-4.421 6.266a25.558 25.558 0 0 0-4.848-.47"/><ellipse cx="33.14" cy="20.309" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.671" ry="2.13"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.714" d="M12.508 40.927c-1.93-.327-4.948-.31-6.04-3.487c-1.067-3.107.438-6.67 3.742-7.045m15.253-4.008a1.467 1.467 0 0 0 1.473-1.472"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.714" d="M28.41 26.387a1.467 1.467 0 0 1-1.474-1.472"/>`),
		g.Group(children),
	)
}