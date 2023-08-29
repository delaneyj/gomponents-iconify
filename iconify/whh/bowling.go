package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bowling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zM191.5 384q-26.5 0-45 18.5t-18.5 45t18.5 45.5t45 19t45.5-19t19-45.5t-19-45t-45.5-18.5zM384 128q-27 0-45.5 18.5t-18.5 45t18.5 45.5t45 19t45.5-19t19-45.5t-18.5-45T384 128zm128 256q-27 0-45.5 18.5t-18.5 45t19 45.5t45 19t45-19t19-45.5t-18.5-45T512 384z"/>`),
		g.Group(children),
	)
}