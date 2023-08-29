package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Robocop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 576h-10q-15 116-49 192q-36 80-97.5 139.5t-127.5 88t-129 28.5t-129-28t-128-87.5T128 768q-34-75-50-192H64q-26 0-45-18.5T0 512V320q0-27 19-45.5T64 256h9q29-110 145-183T483 0t262.5 72.5T888 256h8q26 0 45 18.5t19 45.5v192q0 27-19 45.5T896 576zM183 736q14 32 26 55.5t29.5 50T277 882t43 14q27 0 55.5-16t55-32t49.5-16q22 0 49 16t55.5 32t55.5 16q39 0 76.5-48T782 736q28-63 40-160H142q14 97 41 160zm617-352H160q-13 0-22.5 9.5T128 416t9.5 22.5T160 448h640q13 0 22.5-9.5T832 416t-9.5-22.5T800 384z"/>`),
		g.Group(children),
	)
}