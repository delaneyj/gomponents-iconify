package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.842 5.5l29.375 24.707l-14.712-3.739L11.842 5.5Zm28.593 26.296L6.783 26.615l20.577 9.238l13.075-4.057Zm.171 1.784L19.858 42.5l12.292-.806l8.456-8.114Z"/>`),
		g.Group(children),
	)
}