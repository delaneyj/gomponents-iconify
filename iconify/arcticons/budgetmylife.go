package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Budgetmylife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="12.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.794 29.31a3.138 3.138 0 0 0 3.062 1.49h2.013a4.263 4.263 0 0 0 4.015-3.4h0a2.752 2.752 0 0 0-2.772-3.4h-2.224a2.752 2.752 0 0 1-2.772-3.4h0a4.263 4.263 0 0 1 4.015-3.4h2.013a3.138 3.138 0 0 1 3.062 1.49m-3.963-1.49l.311-1.7m-3.108 17l.311-1.7M24 36.5v8.999M13.194 17.716l-7.779-4.524m29.449 4.626l7.822-4.451"/>`),
		g.Group(children),
	)
}