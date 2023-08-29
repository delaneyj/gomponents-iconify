package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M288 64C64 64 0 160 0 272s80 176 176 176h8.4c24.2 0 46.4-13.7 57.2-35.4l23.2-46.3c4.4-8.8 13.3-14.3 23.2-14.3s18.8 5.5 23.2 14.3l23.2 46.3c10.8 21.7 33 35.4 57.2 35.4h8.4c96 0 176-64 176-176S512 64 288 64zM96 256a64 64 0 1 1 128 0a64 64 0 1 1-128 0zm320-64a64 64 0 1 1 0 128a64 64 0 1 1 0-128z"/>`),
		g.Group(children),
	)
}