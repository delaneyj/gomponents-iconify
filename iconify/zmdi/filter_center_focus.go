package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterCenterFocus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 256v85h85v43H43q-18 0-30.5-12.5T0 341v-85h43zm0-213v85H0V43q0-18 12.5-30.5T43 0h85v43H43zM341 0q18 0 30.5 12.5T384 43v85h-43V43h-85V0h85zm0 341v-85h43v85q0 18-12.5 30.5T341 384h-85v-43h85zM192 128q27 0 45.5 18.5T256 192t-18.5 45.5T192 256t-45.5-18.5T128 192t18.5-45.5T192 128z"/>`),
		g.Group(children),
	)
}