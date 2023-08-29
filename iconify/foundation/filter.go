package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M90 22.292a6.496 6.496 0 0 0-6.258-6.488v-.012H16.848v.018c-.116-.006-.231-.018-.348-.018a6.5 6.5 0 0 0-4.077 11.559l30.14 30.139l-.001 18.599v.154h.015a3.392 3.392 0 0 0 1.713 2.8l-.009.016l7.872 4.545c.066.046.139.079.208.12l.028.016v-.001c.502.29 1.078.469 1.7.469a3.415 3.415 0 0 0 3.416-3.416c0-.09-.02-.175-.026-.263h.026V57.518l30.417-30.416l-.03-.03A6.472 6.472 0 0 0 90 22.292zm-57.751 6.5h.014l.001.015l-.015-.015z"/>`),
		g.Group(children),
	)
}