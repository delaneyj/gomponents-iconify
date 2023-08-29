package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moonreaderpro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37 3a8 8 0 1 0 7.9 8A8 8 0 0 0 37 3Zm4.79 10.3a5.26 5.26 0 1 1-7.2-7.19a.18.18 0 0 1 .25.06a.24.24 0 0 1 0 .12A5.25 5.25 0 0 0 41.57 13a.19.19 0 0 1 .26 0a.18.18 0 0 1-.04.3ZM27.87 25.87L24.9 28l1.1 3.43a.3.3 0 0 1-.18.38a.31.31 0 0 1-.28 0l-2.91-2.21l-2.93 2.17a.3.3 0 0 1-.42-.07a.32.32 0 0 1 0-.27L20.4 28l-3-2.12a.3.3 0 0 1-.06-.42a.28.28 0 0 1 .23-.12h3.62l1.11-3.44a.29.29 0 0 1 .38-.19a.3.3 0 0 1 .19.19L24 25.33h3.6a.29.29 0 0 1 .34.25a.3.3 0 0 1-.12.29Zm7 6.57L33 33.78l.73 2.22a.18.18 0 0 1-.11.24a.21.21 0 0 1-.19 0l-1.85-1.34l-1.85 1.34a.18.18 0 0 1-.26-.05a.23.23 0 0 1 0-.16l.7-2.18l-1.85-1.34a.19.19 0 0 1-.06-.26a.18.18 0 0 1 .17-.09h2.29l.71-2.18a.19.19 0 0 1 .24-.12a.18.18 0 0 1 .12.12l.64 2.18h2.29a.19.19 0 0 1 .22.15a.2.2 0 0 1-.09.2ZM8.4 6.45v35.1a2 2 0 0 0 1.95 2h2.38V4.5h-2.38A2 2 0 0 0 8.4 6.45Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.32 4.5H12.73v39h24.92a2 2 0 0 0 2-2V18.43"/>`),
		g.Group(children),
	)
}