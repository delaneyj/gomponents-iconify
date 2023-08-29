package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Appleflinger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.84 22.29s4.11 1.55 3.15 4c-.23.59-2-.48-2.77-.07M27 39.29c.5.5 1.67 3.21 3.46 3.21s2.46-3.93 3.08-4.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 39.36c-6 0-15.22-.31-15.22-13.38S21.92 9.19 24 9.19S39.22 12.92 39.22 26S30 39.36 24 39.36Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.16 22.29s-4.11 1.55-3.15 4c.23.59 2-.48 2.77-.07M21 39.29c-.5.5-1.67 3.21-3.46 3.21s-2.46-3.93-3.08-4.75m9.09-15.47c2.23 0 7.27-.69 7.27 2.25c0 2.51-4.65 2.26-7.27 2.26m0-4.51c-2.23 0-7.27-.69-7.27 2.25c0 2.51 4.66 2.26 7.27 2.26"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.55 22.28c.66 0 1.43-.07 1.42 1.18a1.31 1.31 0 0 1-1.32 1.18a1.26 1.26 0 0 1-1.33-1c-.06-1.43.4-1.34 1.23-1.36Zm3.83-.04c.66 0 1.48.05 1.33 1.22a1.36 1.36 0 0 1-1.33 1.18a1.25 1.25 0 0 1-1.32-1c-.12-1.64.49-1.37 1.32-1.4Zm-7.47-.01c.67 0 1.33 0 1.33 1.23a1.32 1.32 0 0 1-1.33 1.18a1.24 1.24 0 0 1-1.32-1c0-1.26.49-1.38 1.32-1.41ZM24 9.19a3.46 3.46 0 0 1 2.72-2.85m-6.2-.84A3.77 3.77 0 0 1 24 9.19M17.54 14A3.3 3.3 0 0 1 21 17.17a3.74 3.74 0 0 1-3.58 3.44A3.37 3.37 0 0 1 13.84 17c0-2 1.87-3 3.7-3Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 15.92a.69.69 0 0 1 .64-.14c.6.14 1.38.88 1 1.5a1.3 1.3 0 0 1-1.72.54c-.61-.28-.19-1.85.11-1.9M29.57 14a3.31 3.31 0 0 0-3.47 3.22a3.74 3.74 0 0 0 3.58 3.44A3.37 3.37 0 0 0 33.26 17c0-2-1.87-3-3.69-3Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.07 15.92a.69.69 0 0 0-.64-.14c-.59.14-1.38.88-1 1.5a1.3 1.3 0 0 0 1.72.54c.61-.28.19-1.85-.11-1.9"/>`),
		g.Group(children),
	)
}