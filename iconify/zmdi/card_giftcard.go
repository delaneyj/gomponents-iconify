package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardGiftcard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 88q18 0 30.5 12.5T427 131v234q0 18-12.5 30.5T384 408H43q-18 0-30.5-12.5T0 365V131q0-18 12.5-30.5T43 88h46q-4-11-4-21q0-27 19-45.5T149 3q34 0 54 28l10 15l11-15q19-28 53-28q27 0 45.5 18.5T341 67q0 10-4 21h47zM277.5 45q-8.5 0-15 6.5t-6.5 15t6.5 15t15 6.5t15-6.5t6.5-15t-6.5-15t-15-6.5zm-128 0q-8.5 0-15 6.5t-6.5 15t6.5 15t15 6.5t15-6.5t6.5-15t-6.5-15t-15-6.5zM384 365v-42H43v42h341zm0-106V131H276l44 60l-35 25l-50-69l-22-29l-21 29l-51 69l-34-25l44-60H43v128h341z"/>`),
		g.Group(children),
	)
}