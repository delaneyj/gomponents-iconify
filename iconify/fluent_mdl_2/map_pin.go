package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 2048H0L384 896h512V634q-56-12-103-41t-81-70t-53-94t-19-109q0-66 25-124t68-101t102-69T960 0q66 0 124 25t101 69t69 102t26 124q0 57-19 109t-53 93t-81 71t-103 41v262h512l384 1152zM768 320q0 40 15 75t41 61t61 41t75 15q40 0 75-15t61-41t41-61t15-75q0-40-15-75t-41-61t-61-41t-75-15q-40 0-75 15t-61 41t-41 61t-15 75zM178 1920h1564l-298-896h-420v480q0 9-7 15t-18 10t-21 5t-18 2q-7 0-17-1t-21-5t-18-10t-8-16v-480H476l-298 896z"/>`),
		g.Group(children),
	)
}