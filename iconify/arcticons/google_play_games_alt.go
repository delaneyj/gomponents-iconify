package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlayGamesAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.59 18h0a7.28 7.28 0 0 0-12.24-4.34H18.67a7.29 7.29 0 0 0-5-2h0A7.3 7.3 0 0 0 6.42 18h0L4.51 33.24c-.32 3.89 6.53 4 7.81 1.32l4.17-8.92h15l4.17 8.92c1.28 2.66 8.13 2.57 7.81-1.32Zm-4.68-.47h0a1.42 1.42 0 1 1-1.41 1.42h0a1.42 1.42 0 0 1 1.41-1.46Zm-7.2 1.42a1.42 1.42 0 0 1 1.42-1.42h0a1.42 1.42 0 1 1-1.42 1.42Zm-17.48 3.1v-1.72h-1.72a.87.87 0 0 1-.87-.87v-1.09a.87.87 0 0 1 .87-.88h1.72v-1.71a.87.87 0 0 1 .87-.87h1.09a.87.87 0 0 1 .88.87v1.71h1.71a.87.87 0 0 1 .87.88v1.09a.87.87 0 0 1-.87.87h-1.71v1.72a.87.87 0 0 1-.88.87H13.1a.86.86 0 0 1-.23-.05l-.5-.38a.85.85 0 0 1-.14-.44Zm23.21-.24h0A1.42 1.42 0 1 1 34 20.39h0a1.42 1.42 0 0 1 1.44 1.42ZM34 14.6a1.43 1.43 0 0 1 1.44 1.4h0A1.42 1.42 0 1 1 34 14.6Z"/>`),
		g.Group(children),
	)
}