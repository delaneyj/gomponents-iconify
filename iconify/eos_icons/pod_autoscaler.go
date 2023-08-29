package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodAutoscaler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 11H3.5V8.5L0 12l3.5 3.5V13H5v-2zm19 1l-3.5-3.5V11H19v2h1.5v2.5L24 12zM6 18.752V18h12v1.003A1.998 1.998 0 0 1 16 21H8a1.996 1.996 0 0 1-2-1.997Zm9-14.76a.986.986 0 0 0-.292-.702a.998.998 0 0 0-.706-.29H9.998a.998.998 0 0 0-.92.606a.983.983 0 0 0-.078.386V5h6V3.992Zm3.748 10.258l-1.251-2.75l-1.252-2.75L14.993 6H9L7.75 8.75L6.5 11.5l-1.25 2.75L4 17h16Zm-4.895-2.461a2.007 2.007 0 1 1 .157-.779a2.003 2.003 0 0 1-.157.779Z"/>`),
		g.Group(children),
	)
}