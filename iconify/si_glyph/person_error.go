package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.972 4.071c0 1.664-1.332 4.846-2.976 4.846c-1.645 0-2.979-3.182-2.979-4.846c0-1.665 1.334-3.013 2.979-3.013c1.644.001 2.976 1.348 2.976 3.013zm5.042 6.874l-.996-.996l-1.494 1.493l-1.493-1.493l-.996.996l1.493 1.493l-1.493 1.495l.996.996l1.493-1.495l1.494 1.495l.996-.996l-1.494-1.495l1.494-1.493zm-7.31-.615l-.708.15s-2.904-.283-3.76-1.416c-4.098 0-4.223 5.865-4.223 5.865h9.345l-.63-.65l1.951-1.963l-1.975-1.986z"/>`),
		g.Group(children),
	)
}