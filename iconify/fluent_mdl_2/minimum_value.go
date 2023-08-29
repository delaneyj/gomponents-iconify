package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimumValue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 512h128v896h-128V512zm-256 0h128v896H768V512zm1152 512h-421l162 163l-90 90l-317-317l317-317l90 90l-162 163h421v128zM349 643l317 317l-317 317l-90-90l162-163H0V896h421L259 733l90-90z"/>`),
		g.Group(children),
	)
}