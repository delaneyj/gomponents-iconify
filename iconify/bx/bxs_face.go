package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm0 18c-4.411 0-8-3.589-8-8c0-1.168.258-2.275.709-3.276c.154.09.308.182.456.276c.396.25.791.5 1.286.688c.494.187 1.088.312 1.879.312c.792 0 1.386-.125 1.881-.313s.891-.437 1.287-.687s.792-.5 1.287-.688c.494-.187 1.088-.312 1.88-.312s1.386.125 1.88.313c.495.187.891.437 1.287.687s.792.5 1.287.688c.178.067.374.122.581.171c.191.682.3 1.398.3 2.141c0 4.411-3.589 8-8 8z" fill="currentColor"/><circle cx="8.5" cy="12.5" r="1.5" fill="currentColor"/><circle cx="15.5" cy="12.5" r="1.5" fill="currentColor"/>`),
		g.Group(children),
	)
}