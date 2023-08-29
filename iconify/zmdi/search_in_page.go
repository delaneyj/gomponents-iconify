package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchInPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M203 128q22 0 37.5 15.5T256 181t-15.5 38t-37.5 16t-38-16t-16-38t16-37.5t38-15.5zM384 21q18 0 30.5 12.5T427 64v256q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h341zm-68 303l30-30l-62-62q15-23 15-51q0-40-28-68t-68-28t-68 28t-28 68t28 68t67 28q28 0 51-15z"/>`),
		g.Group(children),
	)
}