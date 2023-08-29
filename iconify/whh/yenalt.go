package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yenalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1023q-104 0-199-40.5t-163.5-109T40.5 710T0 511t40.5-198.5t109-163T313 40.5T512 0t199 40.5t163.5 109t109 163T1024 511t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm302-814q-18-18-43.5-18T727 209L512 371L296 209q-18-18-43.5-18T209 209t-18 43.5t18 43.5l239 180v35h-96q-13 0-22.5 9.5t-9.5 23t9.5 22.5t22.5 9h96v64h-96q-13 0-22.5 9.5T320 671t9.5 22.5T352 703h96v128q0 27 18.5 45.5t45 18.5t45.5-18.5t19-45.5V703h96q13 0 22.5-9.5T704 671t-9.5-22.5T672 639h-96v-64h96q13 0 22.5-9t9.5-22.5t-9.5-23T672 511h-96v-35l238-180q18-18 18-43.5T814 209z"/>`),
		g.Group(children),
	)
}