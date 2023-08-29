package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 .52A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.48A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .52zm0 13.71A6.51 6.51 0 0 1 1.25 8A6.51 6.51 0 0 1 8 1.77A6.51 6.51 0 0 1 14.75 8A6.51 6.51 0 0 1 8 14.23z"/><path fill="currentColor" d="M9 7.38H7A.34.34 0 0 1 6.69 7a.13.13 0 0 1 .01 0a.34.34 0 0 1 .3-.29h1.94a.34.34 0 0 1 .33.29h1.25a1.59 1.59 0 0 0-1.58-1.54h-.32V4.21H7.37v1.25H7A1.59 1.59 0 0 0 5.44 7a1.55 1.55 0 0 0 .35 1A1.59 1.59 0 0 0 7 8.62h2a.34.34 0 0 1 .33.38a.33.33 0 0 1-.33.29H7.08A.33.33 0 0 1 6.75 9H5.49a1.58 1.58 0 0 0 1.58 1.54h.31v1.25h1.24v-1.25H9A1.58 1.58 0 0 0 10.56 9a1.51 1.51 0 0 0-.34-1A1.59 1.59 0 0 0 9 7.38z"/>`),
		g.Group(children),
	)
}