package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Like(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1856 640q39 0 74 15t61 41t42 61t15 75q0 32-10 61l-256 768q-10 29-28 53t-42 42t-52 26t-60 10h-512q-179 0-345-69q-72-29-144-44t-151-15H0V768h417q65 0 122-24t104-70l622-621q25-25 50-39t61-14q33 0 62 12t51 35t34 51t13 62q0 81-18 154t-53 146q-20 43-34 87t-19 93h444zm-256 1024q20 0 37-12t24-32q5-14 18-54t33-96t42-124t46-137t44-134t39-118t27-86t10-39q0-26-19-45t-45-19h-576q0-53 2-98t10-89t22-86t37-91q28-58 42-118t15-126q0-14-9-23t-23-9q-6 0-10 4t-9 9L734 765q-32 32-68 56t-78 41q-80 34-171 34H128v640h320q178 0 345 69q144 59 295 59h512z"/>`),
		g.Group(children),
	)
}