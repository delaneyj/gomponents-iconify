package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Codeigniter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640 1024q22-15 36-36.5t19.5-48.5t7-46t1.5-49q0-38-20.5-77.5t-52-70t-62-52.5t-57.5-36q-54-27-91-68.5T384 448q-17 0-40.5 29.5T320 544q0 18 19.5 39t42.5 38t42 41t19 48q0 49-23 85.5T352 832t-70.5-38t-25.5-90q-18 9-31 23.5t-19.5 28T196 791t-3.5 35t-.5 38q0 114 64 160q-65-12-112.5-36t-75-52.5t-43-69T5 789t-5-85q0-90 93.5-231.5T288 256q5-5 28-22.5t42.5-36T400 151t35-67t13-84q13 0 33.5 12t42 31.5t37 51T576 160q0 32-10 61t-22 47t-22 40.5t-10 43.5q0 40 28 68t68 28q35 0 61.5-17t34.5-47q88 64 140 151t52 169q0 46-8 90t-27 86.5t-47.5 73.5t-73 50.5T640 1024z"/>`),
		g.Group(children),
	)
}