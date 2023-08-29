package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attachment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.404 5.11L9.389 4.096L4.314 9.17a2.152 2.152 0 1 0 3.045 3.044l6.09-6.089a3.588 3.588 0 0 0-5.075-5.074L1.98 7.444l-.014.013c-1.955 1.955-1.955 5.123 0 7.077s5.123 1.954 7.078 0l.013-.014l.001.001l4.365-4.364l-1.015-1.014l-4.365 4.363l-.013.013a3.573 3.573 0 0 1-5.048 0a3.572 3.572 0 0 1 .014-5.06l-.001-.001L9.39 2.065c.839-.84 2.205-.84 3.045 0s.839 2.205 0 3.044l-6.09 6.089a.718.718 0 0 1-1.015-1.014l5.075-5.075z"/>`),
		g.Group(children),
	)
}