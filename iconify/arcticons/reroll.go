package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reroll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.04 5.5l18.5 4.96l4.96 18.5L28.96 42.5l-18.5-4.96l-4.96-18.5L19.04 5.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.25 13.72l-5.22 19.47l19.46-5.22l-14.24-14.25z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.49 27.97l2.05-17.51l-16.29 3.26l14.24 14.25zM21.25 13.72L5.5 19.04l10.53 14.15l5.22-19.47zm-5.22 19.47l12.93 9.31l-18.5-4.96l5.57-4.35zm19.46-5.22L28.96 42.5L42.5 28.96l-7.01-.99zM19.04 5.5l2.21 8.22m.47 14.36V20.8h2.44a2.45 2.45 0 1 1 0 4.89h-2.44m2.44 0l2.42 2.42"/>`),
		g.Group(children),
	)
}