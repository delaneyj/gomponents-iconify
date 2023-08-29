package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryCharging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.56 1.003H9.907V.031H7v.972H5.4c-.766 0-1.384.52-1.384 1.163v12.633c0 .643.618 1.164 1.384 1.164h6.159c.762 0 1.382-.521 1.382-1.164V2.166c0-.644-.62-1.163-1.381-1.163zm.471 13.467c0 .343-.324.62-.72.62H5.632c-.395 0-.719-.277-.719-.62V2.529c0-.343.324-.621.719-.621h5.679c.396 0 .72.278.72.621V14.47z"/><path d="m10 3.006l-3.959 7.683L8.005 9.42l-.989 4.584l3.952-7.821L8.881 7.5L10 3.006z"/></g>`),
		g.Group(children),
	)
}