package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M144 320a80 80 0 1 1 0-160a80 80 0 1 1 0 160zm119.8 0c15.3-22.9 24.2-50.4 24.2-80c0-79.5-64.5-144-144-144S0 160.5 0 240s64.5 144 144 144h352c79.5 0 144-64.5 144-144S575.5 96 496 96s-144 64.5-144 144c0 29.6 8.9 57.1 24.2 80H263.8zM496 160a80 80 0 1 1 0 160a80 80 0 1 1 0-160z"/>`),
		g.Group(children),
	)
}