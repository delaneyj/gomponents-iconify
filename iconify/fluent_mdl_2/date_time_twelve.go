package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DateTimeTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1707 911q78 37 141 93t108 127t68 152t24 168q0 123-47 232t-128 190t-190 128t-232 47q-87 0-168-24t-151-68t-127-107t-94-142H0V171h341V0h171v171h683V0h170v171h342v740zM171 512h1365V341H171v171zm688 1024q-6-42-6-85q0-124 47-233t128-190t190-128t233-47q43 0 85 6V683H171v853h688zm592 341q88 0 165-33t136-92t91-135t34-166q0-88-33-166t-91-136t-136-91t-166-34q-89 0-166 33t-136 92t-91 136t-34 166q0 89 33 166t92 135t136 91t166 34zm85-512h171v171h-342v-341h171v170z"/>`),
		g.Group(children),
	)
}