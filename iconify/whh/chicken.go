package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chicken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 896q-12 0-50-34q-33-11-46 2t-2 46q34 38 34 50q0 26-19 45t-45.5 19t-45-19t-18.5-45q0-12-4-27.5T745 885t-25-52.5t-39-79.5l-72-68q1-2 7-12t11.5-18t12.5-15t14.5-12.5t17-11T684 609l69 72q59 29 79.5 39t52.5 25t47.5 19t27.5 4q26 0 45 19t19 45.5t-19 45t-45 18.5zM608 608q-12 12-21.5 30t-18 26t-24.5 8q-44 0-83.5-10t-65-23.5t-65-32T256 576Q110 527 55 466.5T0 288Q0 169 84.5 84.5T288 0q118 0 178.5 55T576 256q14 41 40.5 92.5t41 95T672 544q0 15-8.5 23.5t-26.5 19t-29 21.5z"/>`),
		g.Group(children),
	)
}