package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HiddenSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.54 38.25l5.11-2.06a15.683 15.683 0 0 0 3.46 2l.77 5.43a1 1 0 0 0 1 .86h8.2a1 1 0 0 0 1-.86l.77-5.43a15.362 15.362 0 0 0 3.46-2l5.11 2.06a1 1 0 0 0 1.25-.44l4.1-7.11a1 1 0 0 0-.25-1.31L39.23 26a16.516 16.516 0 0 0 .14-2a16.516 16.516 0 0 0-.14-2l4.33-3.39m-5.1-8.86l-5.11 2.06a15.679 15.679 0 0 0-3.46-2l-.77-5.43a1 1 0 0 0-1-.86H19.9a1 1 0 0 0-1 .86l-.77 5.43a15.358 15.358 0 0 0-3.46 2L9.54 9.75a1 1 0 0 0-1.25.44l-4.1 7.11a1 1 0 0 0 .25 1.31L8.76 22a16.659 16.659 0 0 0-.14 2a16.516 16.516 0 0 0 .14 2l-4.32 3.39M24 31.18a7.18 7.18 0 1 1 7.17-7.19V24a7.17 7.17 0 0 1-7.16 7.18Z"/>`),
		g.Group(children),
	)
}