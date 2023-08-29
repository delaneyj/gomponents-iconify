package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pizza(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m640 64l-.5-.5l-.5-.5q-2 1-5 2t-5 1l395 958L64 639v1q0 13-9.5 22.5T32 672t-22.5-9.5T0 640v-11q62-229 231-398T630 0h10q13 0 22.5 9t9.5 22.5t-9.5 23T640 64zm127.5 768q26.5 0 45.5-19t19-45.5t-19-45t-45.5-18.5t-45 18.5t-18.5 45t18.5 45.5t45 19zm-512-384q-26.5 0-45 18.5t-18.5 45t18.5 45.5t45 19t45.5-19t19-45.5t-19-45t-45.5-18.5zm192-192q-26.5 0-45 18.5t-18.5 45t18.5 45.5t45 19t45.5-19t19-45.5t-19-45t-45.5-18.5zm0 320q-26.5 0-45 18.5t-18.5 45t18.5 45.5t45 19t45.5-19t19-45.5t-19-45t-45.5-18.5zM576 511.5q0 26.5 18.5 45.5t45.5 19t45.5-19t18.5-45.5t-19-45t-45.5-18.5t-45 18.5t-18.5 45z"/>`),
		g.Group(children),
	)
}