package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.05 16.975q-.5.35-1.025.05t-.525-.9v-8.25q0-.6.525-.9t1.025.05l6.2 4.15q.45.3.45.825t-.45.825l-6.2 4.15Zm10 0q-.5.35-1.025.05t-.525-.9v-8.25q0-.6.525-.9t1.025.05l6.2 4.15q.45.3.45.825t-.45.825l-6.2 4.15ZM4.5 12Zm10 0Zm-10 2.25L7.9 12L4.5 9.75v4.5Zm10 0L17.9 12l-3.4-2.25v4.5Z"/>`),
		g.Group(children),
	)
}