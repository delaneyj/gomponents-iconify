package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cutlery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M640 64v640q0 61-35.5 111T512 885v779q0 52-38 90t-90 38H256q-52 0-90-38t-38-90V885q-57-20-92.5-70T0 704V64q0-26 19-45T64 0t45 19t19 45v416q0 26 19 45t45 19t45-19t19-45V64q0-26 19-45t45-19t45 19t19 45v416q0 26 19 45t45 19t45-19t19-45V64q0-26 19-45t45-19t45 19t19 45zm768 0v1600q0 52-38 90t-90 38h-128q-52 0-90-38t-38-90v-512H800q-13 0-22.5-9.5T768 1120V320q0-132 94-226t226-94h256q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}