package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransitionEffect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 0v896H0V0h128v768h768V347L733 509l-90-90l317-317l317 317l-90 90l-163-162v421h768V0h128zM0 1024h1920v896h-128v-768h-768v421l163-162l90 90l-317 317l-317-317l90-90l163 162v-421H128v768H0v-896z"/>`),
		g.Group(children),
	)
}