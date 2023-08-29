package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rabbit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M767 1025H288q-13 0-22.5-9.5T256 993t9.5-22.5T288 961q14 0 23-9t9-23q0-43-23-97t-71-136.5T160 577q-18-35-41-55.5T76 495t-38-9.5t-28-11T0 449q0-43 44.5-101.5T140 251q-12-88-12-153q0-36 24-66t40-30q11 0 19.5 20.5T223 74q9-33 38-55t44-18q13 4 14.5 32.5T311 98q-12 44-64 142q17 12 24.5 38t17.5 56.5t31 51.5t72.5 37.5t103 26t103 26T671 513q65 64 112.5 163.5T831 833v21q33-21 64-21q27 0 45.5 19t18.5 45q0 30-34.5 61t-79.5 49t-78 18z"/>`),
		g.Group(children),
	)
}