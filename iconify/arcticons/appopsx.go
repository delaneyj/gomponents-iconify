package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Appopsx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.23 26a16.52 16.52 0 0 0 .14-2a16.52 16.52 0 0 0-.14-2l4.33-3.39a1 1 0 0 0 .25-1.31l-4.1-7.11a1 1 0 0 0-1.25-.44l-5.11 2.06a15.68 15.68 0 0 0-3.46-2l-.77-5.43a1.05 1.05 0 0 0-1-.86h-8.2a1 1 0 0 0-1 .86l-.77 5.43a15.36 15.36 0 0 0-3.46 2L9.54 9.75a1 1 0 0 0-1.25.44l-4.1 7.11a1 1 0 0 0 .25 1.31L8.76 22a16.66 16.66 0 0 0-.14 2a16.52 16.52 0 0 0 .14 2l-4.32 3.39a1 1 0 0 0-.25 1.31l4.1 7.11a1 1 0 0 0 1.25.44l5.11-2.06a15.68 15.68 0 0 0 3.46 2l.77 5.43a1 1 0 0 0 1 .86h8.2a1 1 0 0 0 1-.86l.77-5.43a15.36 15.36 0 0 0 3.46-2l5.11 2.06a1 1 0 0 0 1.25-.44l4.1-7.11a1 1 0 0 0-.25-1.31Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.12 14.67l-7.55 3.12v7.66c0 4.12 8.43 8.05 8.43 8.05s8.43-3.93 8.43-8v-7.71l-7.55-3.12a2.33 2.33 0 0 0-1.76 0ZM24 33.5v-19M15.57 24h16.86"/>`),
		g.Group(children),
	)
}