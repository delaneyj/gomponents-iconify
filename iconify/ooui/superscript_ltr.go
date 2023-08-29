package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuperscriptLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 1V0h-.5a1.49 1.49 0 0 0-1 .39a1.49 1.49 0 0 0-1-.39H15v1h.5a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-.5.5H15v1h.5a1.49 1.49 0 0 0 1-.39a1.49 1.49 0 0 0 1 .39h.5V8h-.5a.5.5 0 0 1-.5-.5v-6a.5.5 0 0 1 .5-.5zm-4.32 15h-2.42a.67.67 0 0 1-.46-.15a1.33 1.33 0 0 1-.28-.34l-2.77-4.44a2.65 2.65 0 0 1-.28.69L5 15.51a2.22 2.22 0 0 1-.29.34a.58.58 0 0 1-.42.15H2l4.15-6.19L2.17 4h2.42a.81.81 0 0 1 .41.09a.8.8 0 0 1 .24.26L8 8.59a2.71 2.71 0 0 1 .33-.74L10.6 4.4a.69.69 0 0 1 .6-.4h2.32l-4 5.71z"/>`),
		g.Group(children),
	)
}