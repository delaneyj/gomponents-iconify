package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zM9.51 11.744l6.376-3.668l6.36 3.676l-6.36 3.677l-6.375-3.685zm12.976-.13l3.257-1.9L15.886 4L6 9.714v11.429l9.886 5.714l9.857-5.714l-3.495-2.038l-6.362 3.676l-6.39-3.676v-3.296l6.4 3.696l6.418-3.715v3.315l3.457 2.038V9.714l-3.285 1.9z"/>`),
		g.Group(children),
	)
}