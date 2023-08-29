package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tally(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.661 2.671a.63.63 0 0 1 .595-.49h2c.276 0 .457.219.405.49L5.34 9.545a.63.63 0 0 1-.595.491h-2a.397.397 0 0 1-.405-.49L3.66 2.67Zm5 0a.63.63 0 0 1 .595-.49h2c.276 0 .457.219.405.49L8.071 21.33a.63.63 0 0 1-.594.491h-2a.397.397 0 0 1-.405-.491l3.59-18.658Zm5 0a.63.63 0 0 1 .595-.49h2c.276 0 .457.219.405.49l-3.59 18.658a.63.63 0 0 1-.594.491h-2a.397.397 0 0 1-.405-.491l3.59-18.658Zm5 0a.63.63 0 0 1 .595-.49h2c.276 0 .457.219.405.49L20.34 9.545a.63.63 0 0 1-.595.491h-2a.397.397 0 0 1-.405-.49l1.32-6.876Z"/>`),
		g.Group(children),
	)
}