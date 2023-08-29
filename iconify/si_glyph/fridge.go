package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fridge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 .003c-.553 0-1 .444-1 .992V3.99c0 .548.447.992 1 .992h6c.553 0 1-.444 1-.992V.995a.996.996 0 0 0-1-.992H5zm6.042 2.038H9.958V.958h1.084v1.083zM5 5.99c-.553 0-1 .451-1 1.01v6.894c0 .558.447 1.01 1 1.01L5.021 16H6.98v-1.097h2.041V16h1.959v-1.097c.573 0 1.02-.452 1.02-1.01V6.999c0-.558-.447-1.009-1-1.009H5zm6.042 3.051H9.958V6.875h1.084v2.166z"/>`),
		g.Group(children),
	)
}