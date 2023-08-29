package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.69 2.5a1.09 1.09 0 0 0-1.06.9l-.8 5.7a15.6 15.6 0 0 0-3.64 2.12L8.84 9.06a1.07 1.07 0 0 0-1.31.46L3.23 17a1.09 1.09 0 0 0 .25 1.38L8 21.9a18.53 18.53 0 0 0-.13 2.1A18.21 18.21 0 0 0 8 26.1l-4.52 3.55A1.09 1.09 0 0 0 3.23 31l4.3 7.45a1.06 1.06 0 0 0 1.31.46l5.35-2.16a16.28 16.28 0 0 0 3.64 2.12l.8 5.7a1.09 1.09 0 0 0 1.06.9h8.61a1.09 1.09 0 0 0 1.06-.9l.8-5.7a15.4 15.4 0 0 0 3.64-2.12l5.35 2.16a1.06 1.06 0 0 0 1.31-.46l4.3-7.45a1.09 1.09 0 0 0-.25-1.38L40 26.1a14.77 14.77 0 0 0 0-4.2l4.55-3.55a1.09 1.09 0 0 0 .22-1.35l-4.29-7.48a1.08 1.08 0 0 0-1.32-.46l-5.35 2.16a16.43 16.43 0 0 0-3.63-2.12l-.81-5.7a1.11 1.11 0 0 0-1.07-.9ZM24.13 24H33a8.88 8.88 0 0 1-8.83 9.09a9.09 9.09 0 1 1 4.5-17"/>`),
		g.Group(children),
	)
}