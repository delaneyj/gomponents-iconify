package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microgsettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.23 26a16.52 16.52 0 0 0 .14-2a16.52 16.52 0 0 0-.14-2l4.33-3.39a1 1 0 0 0 .25-1.31l-4.1-7.11a1 1 0 0 0-1.25-.44l-5.11 2.06a15.68 15.68 0 0 0-3.46-2l-.77-5.43a1 1 0 0 0-1-.86H19.9a1 1 0 0 0-1 .86l-.77 5.43a15.36 15.36 0 0 0-3.46 2L9.54 9.75a1 1 0 0 0-1.25.44l-4.1 7.11a1 1 0 0 0 .25 1.31L8.76 22a16.66 16.66 0 0 0-.14 2a16.52 16.52 0 0 0 .14 2l-4.32 3.39a1 1 0 0 0-.25 1.31l4.1 7.11a1 1 0 0 0 1.25.44l5.11-2.06a15.68 15.68 0 0 0 3.46 2l.77 5.43a1 1 0 0 0 1 .86h8.2a1 1 0 0 0 1-.86l.77-5.43a15.36 15.36 0 0 0 3.46-2l5.11 2.06a1 1 0 0 0 1.25-.44l4.1-7.11a1 1 0 0 0-.25-1.31Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".93" d="m17.26 22.34l-4.72-1.18m6.79-7.92l1.93 4.58m-2.34 16.73l2.1-4.48m7.6-10.86a6.84 6.84 0 0 0-4.75-1.89a6.68 6.68 0 1 0 0 13.36a6.9 6.9 0 0 0 4.77-1.9l3.45 3.45a11.66 11.66 0 1 1-8.6-19.89h0a11.79 11.79 0 0 1 1.88.1a11.61 11.61 0 0 1 6.72 3.31Z"/>`),
		g.Group(children),
	)
}