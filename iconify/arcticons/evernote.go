package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Evernote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.17 4.67a6.15 6.15 0 0 1 4.29.13a4.39 4.39 0 0 1 2.23 2.5s8.49.43 11 2.87s2.4 8.74 2.58 11.5c.21 3.16.63 11-.8 14.55c-1.32 3.21-3.83 6.65-7.67 7.07a6.64 6.64 0 0 1-6.76-3a5.28 5.28 0 0 1 0-5.18a4.06 4.06 0 0 1 3.19-1.85c.72 0 1.69.27 1.94.95a1.06 1.06 0 0 1-.6 1.4c-.49.26-1 .07-1.55.58a2 2 0 0 0-.43 2a2.2 2.2 0 0 0 1.59 1.47a5.08 5.08 0 0 0 4.48-2a3.94 3.94 0 0 0 .17-4a5.43 5.43 0 0 0-3.45-2.85a15.66 15.66 0 0 1-4.1-1.37a9.64 9.64 0 0 1-2.19-2.76S21.92 31 20.2 32c-1.9 1.11-4.56.75-6.6-.06a12.27 12.27 0 0 1-6-5.61c-1.07-2-1.1-4.87-1-8.87c0-2.29 10.29-12 12.63-12.86Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.22 22.33c-.77-.07.67-1.76 1.61-1.73c1.6 0 3.42 2.15 2.55 2.08ZM19.17 4.67c-3 2.72-1.69 4-2 9.08a1.94 1.94 0 0 1-1.91 1.94c-4.51.44-7.42-.77-8.7 1.78"/>`),
		g.Group(children),
	)
}