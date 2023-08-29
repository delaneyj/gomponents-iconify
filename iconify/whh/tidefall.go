package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tidefall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 591q-53-18-118-15t-125.5 20.5t-130 39t-132 40.5T258 701.5T128 696q-54-11-91-35.5T0 609q0-29 38-43t90-4q43 19 103 16.5T352 558t133-39.5t139-41t139-26t133 5.5q54 11 91 35.5t37 51.5q0 30-39.5 46t-88.5 1zM536 312q-10 8-24 8t-23-8L327 166q-17-16 8-38h113V32q0-13 9.5-22.5T480 0h64q13 0 22.5 9.5T576 32v96h113q25 23 9 38zM128 882q43 19 103 16t121-20.5t133-39t139-41t139-26t133 5.5q54 11 91 35.5t37 51.5q0 30-39.5 46t-88.5 1q-53-18-118-15t-125.5 20.5t-130 39t-132 40.5t-132.5 25.5t-130-5.5q-54-11-91-35.5T0 928t38-42t90-4z"/>`),
		g.Group(children),
	)
}