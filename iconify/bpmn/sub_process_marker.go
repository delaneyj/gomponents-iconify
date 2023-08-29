package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubProcessMarker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M300 300v1400h1400V300zm88 88h1224v1224H388zm522 212v310H600v180h310v310h180v-310h310V910h-310V600z"/>`),
		g.Group(children),
	)
}