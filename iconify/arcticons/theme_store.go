package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThemeStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.4 26.829a3.618 3.618 0 0 0 .699-4.143c-1.356-2.794-3.577-6.44-6.936-9.857a33.503 33.503 0 0 0-9.932-7.004c-1.366-.632-2.98-.32-4.045.744L9.172 18.583a3.62 3.62 0 0 0 0 5.12l5.047 5.047l-7.545 7.544c-1.36 1.36-1.56 3.598-.273 5.027a3.552 3.552 0 0 0 5.155.139l7.686-7.687l5.048 5.047a3.62 3.62 0 0 0 5.119 0L41.4 26.83ZM13.489 14.265l20.238 20.237"/>`),
		g.Group(children),
	)
}