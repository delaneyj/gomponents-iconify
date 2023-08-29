package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowflake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M36 7v58M11.237 50.502l49.526-29.004m0 29.004L11.237 21.498m9.806 5.742V16m4.914-2.292L36 18.458M21.043 27.24L11 31.991m35.043-18.283L36 18.458M25.957 58.209L36 53.458m10.043 4.751L36 53.458m-14.923-8.716l-10.043-4.75M61 31.991L50.957 27.24m10.011 12.752l-10.043 4.75m0 11.241V44.742M21.077 55.983V44.742m29.88-17.502V16"/>`),
		g.Group(children),
	)
}