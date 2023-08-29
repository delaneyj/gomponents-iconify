package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M440 368a24.028 24.028 0 0 0 24-24V104a24.028 24.028 0 0 0-24-24H72a24.028 24.028 0 0 0-24 24v240a24.028 24.028 0 0 0 24 24ZM80 112h352v224H80ZM16 400h480v32H16z"/>`),
		g.Group(children),
	)
}