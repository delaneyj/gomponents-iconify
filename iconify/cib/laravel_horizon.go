package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaravelHorizon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.313 4.688c-6.245-6.25-16.375-6.25-22.625-.005A15.966 15.966 0 0 0 0 15.996a15.958 15.958 0 0 0 5.599 12.177h.016c6.349 5.417 15.797 5.042 21.698-.859c6.25-6.245 6.25-16.375 0-22.625zm-5.98 13.973c-5.333 0-5.333-5.333-10.667-5.333c-3.333 0-4.589 2.089-6.354 3.656h-.005C3.76 10.526 8.557 4.849 15.015 4.302c6.458-.542 12.13 4.25 12.677 10.708c-1.771 1.568-3.026 3.651-6.359 3.651z"/>`),
		g.Group(children),
	)
}