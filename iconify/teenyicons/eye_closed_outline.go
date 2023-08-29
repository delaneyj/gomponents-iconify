package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeClosedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 9C5.186 9 3.561 7.848 2.497 6.666a9.368 9.368 0 0 1-1.449-2.164a5.065 5.065 0 0 1-.08-.18l-.004-.007v-.001L.5 4.5l-.464.186v.002l.003.004a2.107 2.107 0 0 0 .026.063l.078.173a10.367 10.367 0 0 0 1.61 2.406C2.94 8.652 4.814 10 7.5 10V9Zm7-4.5a68.887 68.887 0 0 1-.464-.186l-.003.008l-.015.035l-.066.145a9.37 9.37 0 0 1-1.449 2.164C11.44 7.848 9.814 9 7.5 9v1c2.686 0 4.561-1.348 5.747-2.666a10.365 10.365 0 0 0 1.61-2.406a6.164 6.164 0 0 0 .104-.236l.002-.004v-.001h.001L14.5 4.5ZM8 12V9.5H7V12h1Zm-6.646-1.646l2-2l-.708-.708l-2 2l.708.708Zm10.292-2l2 2l.708-.708l-2-2l-.708.708Z"/>`),
		g.Group(children),
	)
}