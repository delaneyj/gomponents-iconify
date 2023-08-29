package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lanyrd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5zm-1.65 12.012l-5.444 1.781c-1.244.406-1.369.341-1.931-1.4L4.1 8.134c-.328-1.009-1.328-3.728-1.497-4.25c-.313-.969-.313-1.022 1.516-1.616c1.431-.469 1.491-.453 2.009 1.163c.419 1.3.688 2.35 1.119 3.678l1.172 3.625l3.744-1.225c.738-.244.984-.231 1.194.678l.15.688c.175.797-.228 1-.656 1.137z"/>`),
		g.Group(children),
	)
}