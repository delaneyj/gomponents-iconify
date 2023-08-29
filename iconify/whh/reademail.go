package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reademail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1017.713 989l-350-349l349-349q8 14 8 29v640q0 15-7 29zm-120-669h-1v1l-256 256v-1h-256v1l-256-256v-1h-1l-63-63V64q0-27 19-45.5t45-18.5h768q27 0 45.5 18.5t18.5 45.5v193zm-540 320l-350 349q-7-14-7-29V320q0-15 8-29zm220 0l384 384h-898l384-384h130z"/>`),
		g.Group(children),
	)
}