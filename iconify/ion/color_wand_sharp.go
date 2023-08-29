package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorWandSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m133.441 200.647l67.197-67.196l78.142 78.142l-67.196 67.197zM301.41 234.21l-67.19 67.2L412 480l68-68l-178.59-177.79zM32 176h80v32H32zm35.624-85.75L90.25 67.622l56.569 56.569l-22.628 22.627zM176 32h32v80h-32zm61.32 92.195l56.569-56.569l22.627 22.627l-56.568 56.569zM67.62 293.886l56.569-56.568l22.627 22.627l-56.568 56.569z"/>`),
		g.Group(children),
	)
}