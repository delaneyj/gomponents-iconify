package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.001 2.5v19l-18-2v-15l18-2Zm-2 10.499l-7 .001v5.487l7 .779v-6.267Zm-14 4.71l5 .556V13l-5-.001v4.71Zm14-6.71V4.735l-7 .777V11l7-.001Zm-9-5.265l-5 .556V11l5 .001V5.734Z"/>`),
		g.Group(children),
	)
}