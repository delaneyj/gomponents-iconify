package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3 1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H3zm2.75.5h3.51a.25.25 0 1 1 0 .5H5.75a.25.25 0 1 1 0-.5zM3.5 3H7v4H3.5a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5zM8 3h3.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5H8V3zM5 8a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm5 0a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm.445 3.994a.504.504 0 0 0-.425.676l.17.33H4.81l.13-.27a.5.5 0 0 0-.91-.41l-1 2a.487.487 0 0 0-.03.18a.5.5 0 0 0 .5.5a.49.49 0 0 0 .43-.26v-.05H4l.31-.69h6.38l.31.69v.05a.49.49 0 0 0 .43.26a.5.5 0 0 0 .5-.5a.49.49 0 0 0 0-.24l-1-2a.5.5 0 0 0-.485-.266z"/>`),
		g.Group(children),
	)
}