package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libremsocial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.77 26l-3.65 1.39v11a1.31 1.31 0 0 1-2.61.25a1 1 0 0 1 0-.25v-.62a45.34 45.34 0 0 0-17.82-6.81l.39 1.21a1 1 0 0 1-.64 1.18a.84.84 0 0 1-.27 0h-.62l1.29 4.32a1 1 0 0 1-.65 1.18a1.12 1.12 0 0 1-.27 0h-4.05a.89.89 0 0 1-.87-.65l-2.31-7.74h-2.2a2 2 0 0 1-2-2v-.92a1.9 1.9 0 0 1-1.95-1.84v-4.25a1.91 1.91 0 0 1 1.89-1.92h0v-.89a2 2 0 0 1 2-2h5.83c8 0 16.79-3.47 22.28-7.15v-.68a1.32 1.32 0 0 1 2.62 0v11l3.65 1.39a1.14 1.14 0 0 1 .72 1.06v2.65a1.14 1.14 0 0 1-.76 1.09Zm-3.61-6.18v7.54"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.71 30.46a55.67 55.67 0 0 1 8 .42"/>`),
		g.Group(children),
	)
}