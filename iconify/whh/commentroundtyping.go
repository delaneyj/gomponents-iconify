package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commentroundtyping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-161 0-293-93q-35 43-93 68T0 1024q25 0 43.5-76.5T64 759Q0 644 0 512q0-139 68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zM320.5 448q-26.5 0-45.5 18.5T256 512t19 45.5t45.5 18.5t45-19t18.5-45t-18.5-45t-45-19zm192 0q-26.5 0-45.5 18.5t-19 45t19 45.5t45.5 19t45-18.5T576 512t-18.5-45.5t-45-18.5zm192 0q-26.5 0-45.5 19t-19 45t19 45t45.5 19t45-19t18.5-45t-18.5-45t-45-19z"/>`),
		g.Group(children),
	)
}