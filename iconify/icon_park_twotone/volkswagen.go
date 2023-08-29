package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volkswagen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVolkswagen0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m7 14l9 23l6-11h4l6 11l9-23"/><path stroke-linecap="round" stroke-linejoin="round" d="m16 6l6 13h4l6-13"/><path stroke-linecap="round" d="M44 24a19.952 19.952 0 0 0-6.77-15A19.924 19.924 0 0 0 24 4a19.924 19.924 0 0 0-13.23 5A19.952 19.952 0 0 0 4 24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVolkswagen0)"/>`),
		g.Group(children),
	)
}