package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barchartasc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M993 1025H801q-13 0-22.5-9.5T769 993V289q0-14 9.5-23t22.5-9h192q13 0 22.5 9t9.5 23v704q0 13-9.5 22.5T993 1025zm-320 0H481q-14 0-23-9.5t-9-22.5V545q0-14 9-23t23-9h192q13 0 22.5 9t9.5 23v448q0 13-9.5 22.5T673 1025zm0-704L561 209L117 564q-23 15-52.5 12.5t-48-21t-15-42.5T27 474l445-355l-87-87q0-13 9-22.5T417 0h256q13 0 22.5 9.5T705 32v257q0 13-9.5 22.5T673 321zM160 705h193q13 0 22.5 9.5T385 737v256q0 13-9.5 22.5T353 1025H160q-13 0-22.5-9.5T128 993V737q0-13 9.5-22.5T160 705z"/>`),
		g.Group(children),
	)
}