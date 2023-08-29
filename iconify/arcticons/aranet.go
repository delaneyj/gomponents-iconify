package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aranet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.286L4.5 40.714h39L24 7.286z"/><circle cx="17.5" cy="25.857" r="1.393" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="14.714" r=".929" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.643" cy="22.143" r=".929" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="32.357" cy="29.571" r=".929" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37" cy="37" r=".929" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11" cy="37" r=".929" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.794 27.064l-5.33 9.137m12.072-20.688l-5.34 9.146m.697.734l8.868-2.953m3.695 6.909l-12.61-3.158m17.355 10.354L18.707 26.544"/>`),
		g.Group(children),
	)
}