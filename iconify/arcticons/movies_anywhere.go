package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoviesAnywhere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.073 25.33l-6.81 14.014c-.37.765-.675.698-.683-.15c-.093-9.825-.577-19.695 2.642-29.14a.859.859 0 0 1 1.575-.312c6.443 7.57 8.658 22.178 13.197 24.838a3.965 3.965 0 0 0 2.995.086c5.395-2.458 7.757-17.84 13.76-25.098a.839.839 0 0 1 1.55.311c3.467 10.383 3.485 19.92 2.937 29.317c-.05.846-.407.917-.794.16L35.281 25.33"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.1 18.998s4.16-8.458 7.774-10.536a3.962 3.962 0 0 1 2.98-.1c4.104 1.862 8.13 10.38 8.13 10.38"/>`),
		g.Group(children),
	)
}