package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallzShooter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="27.317" cy="20.082" r="3.794" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.9 28.866l2.618-4.21m.374 4.96l2.645-4.254m-7.673 1.251l2.558-4.115M5.628 31.223a31.2 31.2 0 0 1 22.846 11.256M16.863 27.55l8.916 5.148l-1.469 2.587a31.993 31.993 0 0 0-9.004-5.018Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.618 11.655h7.992V5.5m20.78 0v6.155h7.895M15.973 5.5v6.155h16.054V5.5"/>`),
		g.Group(children),
	)
}