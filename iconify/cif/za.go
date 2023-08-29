package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Za(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<defs><path id="cifZa0" d="M0 0h300v200H0z"/><path id="cifZa1" d="m0 0l150 100L0 200"/></defs><g fill="none" fill-rule="evenodd"><mask id="cifZa2" fill="#fff"><use href="#cifZa0"/></mask><g mask="url(#cifZa2)"><path fill="#002395" fill-rule="nonzero" d="M0 0v200h300V0z"/><path fill="#DE3831" fill-rule="nonzero" d="M0 0v100h300V0z"/><path fill="#000" fill-rule="nonzero" stroke="#FFF" stroke-width="66.667" d="m0 0l150 100L0 200m146-100h154"/><mask id="cifZa3" fill="#fff"><use href="#cifZa1"/></mask><path fill="#000" fill-rule="nonzero" stroke="#FFB612" stroke-width="60" d="m0 0l150 100L0 200" mask="url(#cifZa3)"/><path stroke="#007A4D" stroke-width="40" d="m0 0l150 100L0 200m150-100h150"/></g></g>`),
		g.Group(children),
	)
}