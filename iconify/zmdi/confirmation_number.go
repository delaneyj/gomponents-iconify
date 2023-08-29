package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConfirmationNumber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 149q-18 0-30.5 12.5T384 192t12.5 30.5T427 235v85q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320v-85q18 0 30.5-12.5T43 192t-12.5-30.5T0 149V64q0-18 12.5-30.5T43 21h341q18 0 30.5 12.5T427 64v85zM235 309v-42h-43v42h43zm0-96v-42h-43v42h43zm0-96V75h-43v42h43z"/>`),
		g.Group(children),
	)
}