package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Route(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.016 4.502A2.975 2.975 0 0 0 3.038 7.48c0 2.233 2.978 5.53 2.978 5.53s2.978-3.297 2.978-5.53a2.975 2.975 0 0 0-2.978-2.978Zm0 4.041A1.063 1.063 0 1 1 7.079 7.48a1.064 1.064 0 0 1-1.063 1.063Zm15.008 2.753v-4.3a4.962 4.962 0 0 0-.204-1.333a4.996 4.996 0 0 0-9.796 1.216v.248l-.01.87v9.952h-.004v.041a2 2 0 0 1-4 0c0-.012.003-.024.004-.037H7.01V16.01h-2v2h.002a3.998 3.998 0 0 0 7.996-.005h.002v-.982h.005V8.997l.01-1.87V6.88a3.001 3.001 0 0 1 6 .123v4.275a1.999 1.999 0 1 0 2 .018Z"/>`),
		g.Group(children),
	)
}