package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VimVisualModeSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1.75C0 .784.784 0 1.75 0h12.5C15.216 0 16 .784 16 1.75v12.5A1.75 1.75 0 0 1 14.25 16H1.75A1.75 1.75 0 0 1 0 14.25V1.75Zm1.75-.25a.25.25 0 0 0-.25.25v12.5c0 .138.112.25.25.25h12.5a.25.25 0 0 0 .25-.25V1.75a.25.25 0 0 0-.25-.25H1.75Zm10.005 2.295c.39.14.591.57.45.96l-2.376 6.58c-.542 1.501-2.647 1.552-3.261.078l-2.76-6.625a.75.75 0 0 1 1.384-.576l2.76 6.624a.25.25 0 0 0 .467-.01l2.376-6.58a.75.75 0 0 1 .96-.451Z"/>`),
		g.Group(children),
	)
}