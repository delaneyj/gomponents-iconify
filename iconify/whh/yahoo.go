package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yahoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M991.998 256h-119l-297 328v312h96q13 0 22.5 9.5t9.5 22.5v64q0 13-9.5 22.5t-22.5 9.5h-384q-13 0-22.5-9.5t-9.5-22.5v-64q0-13 9.5-22.5t22.5-9.5h96V591l-262-461v-2h-90q-13 0-22.5-9.5T-.002 96V32q0-13 9.5-22.5t22.5-9.5h384q13 0 22.5 9.5t9.5 22.5v64q0 13-9.5 22.5t-22.5 9.5h-126l212 367l218-239h-112q-13 0-22.5-9.5t-9.5-22.5v-64q0-13 9.5-22.5t22.5-9.5h384q13 0 22.5 9.5t9.5 22.5v64q0 13-9.5 22.5t-22.5 9.5z"/>`),
		g.Group(children),
	)
}