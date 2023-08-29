package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1408 1536l128-128v384H0V512h128v1152h1280v-128zm-128-512q-128 0-250 25t-237 75t-217 121t-192 163v-128q0-124 32-238t90-214t140-181t181-140t214-91t239-32V0l704 704l-704 704v-384zm101-512q-56 0-105 1t-97 6t-96 18t-102 35q-87 36-161 92T687 791t-99 155t-60 176q168-112 359-169t393-57h128v203l395-395l-395-395v203h-27z"/>`),
		g.Group(children),
	)
}