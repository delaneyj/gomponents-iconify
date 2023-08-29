package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eruditetranslator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 13.003h11.652l5.082-4.949l5.085 4.95h3.073v5.233H43.5v16.757H31.853l-5.06 4.952l-5.081-4.952h-3.076V29.76H4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 19.415a1.748 1.748 0 0 1 1.719 1.775h0v2.195A1.748 1.748 0 0 1 24 25.16h0a1.748 1.748 0 0 1-1.719-1.775h0V21.19A1.748 1.748 0 0 1 24 19.415Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.77 24.23a3.28 3.28 0 0 0 6.46 0M24 26.938v1.647m8.348-4.512h9.038m-4.544-1.969v1.969m2.817 0c0 2.719-4.047 6.343-7.244 7.03"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.9 24.35c.29 2.108 4.226 6.316 7.155 6.754M7.723 26.678l3.526-10.636m3.379 10.668l-3.379-10.668m2.249 7.099H8.895m7.257-10.138H26.32m-4.608 21.991h10.141"/><ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.198" ry="6.206"/>`),
		g.Group(children),
	)
}