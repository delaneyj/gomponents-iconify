package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrCardboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M576 64H64C28.7 64 0 92.7 0 128v256c0 35.3 28.7 64 64 64h120.4c24.2 0 46.4-13.7 57.2-35.4l32-64c8.8-17.5 26.7-28.6 46.3-28.6s37.5 11.1 46.3 28.6l32 64c10.8 21.7 33 35.4 57.2 35.4H576c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64zM96 240a64 64 0 1 1 128 0a64 64 0 1 1-128 0zm384-64a64 64 0 1 1 0 128a64 64 0 1 1 0-128z"/>`),
		g.Group(children),
	)
}