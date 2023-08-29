package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TranslateOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 2v2.5H2v2h7.46a14.98 14.98 0 0 1-2.085 3.665A15.015 15.015 0 0 1 5.91 7.868l-.459-.889l-1.777.917l.458.889a17.022 17.022 0 0 0 1.894 2.9a15.063 15.063 0 0 1-3.028 2.309l-.865.5l1.001 1.732l.866-.501a17.069 17.069 0 0 0 3.374-2.56a17.06 17.06 0 0 0 3.374 2.56l.866.501l1.001-1.731l-.865-.5a15.062 15.062 0 0 1-3.029-2.31A16.97 16.97 0 0 0 11.59 6.5H13v-2H8.5V2h-2Zm10 7.164l-5.832 12.312l1.808.856L13.58 20h5.84l1.104 2.332l1.808-.856L16.5 9.164ZM18.472 18h-3.944l1.972-4.164L18.472 18Z"/>`),
		g.Group(children),
	)
}