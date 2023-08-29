package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextFrameTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 3.75a.75.75 0 0 0-1.5 0v16.5a.75.75 0 0 0 1.5 0V3.75Zm6.24-.44C10.577 2.513 9 3.344 9 4.752v14.495c0 1.413 1.589 2.244 2.75 1.437l10.498-7.302a1.75 1.75 0 0 0-.01-2.88L11.739 3.309ZM10.5 4.752a.25.25 0 0 1 .391-.206L21.39 11.74a.25.25 0 0 1 .001.412l-10.498 7.301a.25.25 0 0 1-.393-.205V4.753Z"/>`),
		g.Group(children),
	)
}