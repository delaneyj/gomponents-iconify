package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 0h128l39 43h68q18 0 30.5 12.5T427 85v256q0 18-12.5 30.5T384 384H43q-18 0-30.5-12.5T0 341V85q0-17 12.5-29.5T43 43h67zm64 320q44 0 75.5-31.5T320 213t-31.5-75t-75.5-31t-75 31t-31 75t31 75.5t75 31.5zm0-21l-26-59l-59-27l59-26l26-59l27 59l59 26l-59 27z"/>`),
		g.Group(children),
	)
}