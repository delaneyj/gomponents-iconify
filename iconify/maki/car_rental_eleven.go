package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarRentalEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.025 7H8.84l-.278-1.67a.988.988 0 0 0-.979-.83H3.417c-.488 0-.9.35-.979.83L2.16 7H1.98A.98.98 0 0 0 1 7.98V10h1a1 1 0 0 0 2 0h3a1 1 0 0 0 2 0h1V7.975A.975.975 0 0 0 9.025 7zM3.25 9a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5zm-.08-2l.247-1.5l4.158-.007L7.826 7H3.171zm4.58 2a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5z" fill="currentColor"/><path d="M9.497 0h-1.5A.997.997 0 0 0 7 .997V1H2.5l-1 1l1 1l1-1l1 1l1-1l1 1H7c.003.553.45 1 1.004 1h1.494A.502.502 0 0 0 10 3.498V.503A.503.503 0 0 0 9.497 0zM9 2.5a.5.5 0 0 1-1 0v-1a.5.5 0 1 1 1 0v1z" fill="currentColor"/>`),
		g.Group(children),
	)
}