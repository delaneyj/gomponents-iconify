package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wacom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1004 119L866 257h29q53 0 90.5 37.5T1023 385v511q0 53-37.5 90.5T895 1024H128q-53 0-90.5-37.5T0 896V385q0-53 37.5-90.5T128 257h541L905 20q21-20 49.5-20t49 20.5t20.5 49t-20 49.5zM192 353q0-14-9.5-23t-22.5-9H96q-13 0-22.5 9T64 353v192q0 13 9.5 22t22.5 9h64q13 0 22.5-9t9.5-22V353zm0 383q0-13-9.5-22.5T160 704H96q-13 0-22.5 9.5T64 736v192q0 13 9.5 22.5T96 960h64q13 0 22.5-9.5T192 928V736zm128-415q-27 0-45.5 18.5T256 385v511q0 27 18.5 45.5T320 960h575q26 0 45-18.5t19-45.5V385q0-27-18.5-45.5T895 321h-93L516 606q-12 12-45 21t-60 11l-28 3q1-12 3-30.5t11-54.5t21-48l187-187H320z"/>`),
		g.Group(children),
	)
}