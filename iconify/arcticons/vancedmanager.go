package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vancedmanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.23 26a16.52 16.52 0 0 0 .14-2a16.52 16.52 0 0 0-.14-2l4.33-3.39a1 1 0 0 0 .25-1.31l-4.1-7.11a1 1 0 0 0-1.25-.44l-5.11 2.06a15.68 15.68 0 0 0-3.46-2l-.77-5.43a1.05 1.05 0 0 0-1-.86h-8.2a1 1 0 0 0-1 .86l-.77 5.43a15.36 15.36 0 0 0-3.46 2L9.54 9.75a1 1 0 0 0-1.25.44l-4.1 7.11a1 1 0 0 0 .25 1.31L8.76 22a16.66 16.66 0 0 0-.14 2a16.52 16.52 0 0 0 .14 2l-4.32 3.39a1 1 0 0 0-.25 1.31l4.1 7.11a1 1 0 0 0 1.25.44l5.11-2.06a15.68 15.68 0 0 0 3.46 2l.77 5.43a1 1 0 0 0 1 .86h8.2a1 1 0 0 0 1-.86l.77-5.43a15.36 15.36 0 0 0 3.46-2l5.11 2.06a1 1 0 0 0 1.25-.44l4.1-7.11a1 1 0 0 0-.25-1.31Zm-19.32-4.65a.52.52 0 0 1 .8-.44l4 2.6a.52.52 0 0 1 0 .87l-4 2.6a.52.52 0 0 1-.8-.44Zm.38-4.44l9.5 6.22a1 1 0 0 1 0 1.63L20.29 31"/>`),
		g.Group(children),
	)
}