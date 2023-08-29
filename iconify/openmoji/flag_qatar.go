package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagQatar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#781e32" d="M5 17h62v38H5z"/><path fill="#fff" d="M5 17h14l7 2.5l-6.5 2.4l6.5 2.3l-6.5 2.4l6.5 2.3l-6.5 2.4l6.5 2.4l-6.5 2.3l6.5 2.4l-6.5 2.3l6.5 2.4l-6.5 2.4l6.5 2.3l-6.5 2.4l6.5 2.3l-7 2.5H5V17z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}