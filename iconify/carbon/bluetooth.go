package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 30a1.06 1.06 0 0 1-.42-.09A1 1 0 0 1 15 29V18.41L8.41 25L7 23.59L14.59 16L7 8.41L8.41 7L15 13.59V3a1 1 0 0 1 .58-.91a1 1 0 0 1 1.07.15l7 6A1 1 0 0 1 24 9a1 1 0 0 1-.29.75L17.41 16l6.3 6.29A1 1 0 0 1 24 23a1 1 0 0 1-.35.72l-7 6A1 1 0 0 1 16 30Zm1-11.59v8.42l4.53-3.89Zm0-13.24v8.42l4.53-4.53Z"/>`),
		g.Group(children),
	)
}