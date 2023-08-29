package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mayanpyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 768H704L576 0h96q13 0 22.5 9.5T704 32v64q0 13-9.5 22.5T672 128h64q13 0 22.5 9.5T768 160v64q0 13-9.5 22.5T736 256h64q13 0 22.5 9t9.5 23v64q0 13-9.5 22.5T800 384h64q13 0 22.5 9t9.5 23v64q0 13-9.5 22.5T864 512h64q13 0 22.5 9t9.5 23v64q0 13-9.5 22.5T928 640h64q13 0 22.5 9t9.5 23v64q0 13-9.5 22.5T992 768zm-608 0L448 0h128l64 768H384zm-352 0q-13 0-22.5-9.5T0 736v-64q0-13 9.5-22.5T32 640h64q-13 0-22.5-9.5T64 608v-64q0-13 9.5-22.5T96 512h64q-13 0-22.5-9.5T128 480v-64q0-13 9.5-22.5T160 384h64q-13 0-22.5-9.5T192 352v-64q0-13 9.5-22.5T224 256h64q-13 0-22.5-9.5T256 224v-64q0-13 9.5-22.5T288 128h64q-13 0-22.5-9.5T320 96V32q0-13 9.5-22.5T352 0h96L320 768H32z"/>`),
		g.Group(children),
	)
}