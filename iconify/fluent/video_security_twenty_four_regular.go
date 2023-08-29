package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoSecurityTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 3A3.25 3.25 0 0 0 2 6.25v6.5A3.25 3.25 0 0 0 5.25 16h2.243a1.656 1.656 0 0 1-1.65 1.5H3.25a.75.75 0 0 0-.75.75v3c0 .414.336.75.75.75h3A5.75 5.75 0 0 0 12 16.25V16h1.75a3.25 3.25 0 0 0 3.235-2.934l3.88 2.327A.75.75 0 0 0 22 14.75V4.25a.75.75 0 0 0-1.136-.643l-3.88 2.327A3.25 3.25 0 0 0 13.75 3h-8.5ZM17 7.675l3.5-2.1v7.85l-3.5-2.1v-3.65ZM6.25 20.5H4V19h1.844c1.69 0 3.07-1.33 3.152-3H10.5v.25a4.25 4.25 0 0 1-4.25 4.25ZM3.5 6.25c0-.966.784-1.75 1.75-1.75h8.5c.966 0 1.75.784 1.75 1.75v6.5a1.75 1.75 0 0 1-1.75 1.75h-8.5a1.75 1.75 0 0 1-1.75-1.75v-6.5Z"/>`),
		g.Group(children),
	)
}