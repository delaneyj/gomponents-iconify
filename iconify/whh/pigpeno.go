package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigpeno(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M384.784 512q0-53 37.5-90.5t90.5-37.5t90.5 37.5t37.5 90.5t-37.5 90.5t-90.5 37.5t-90.5-37.5t-37.5-90.5zm544-320h-736v640h736q40 0 68 28t28 68t-28 68t-68 28h-832q-40 0-68-28t-28-68V96q0-40 28-68t68-28h832q40 0 68 28t28 68t-28 68t-68 28z"/>`),
		g.Group(children),
	)
}