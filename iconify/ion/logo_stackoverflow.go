package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoStackoverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M392 440V320h40v160H64V320h40v120Z"/><path fill="currentColor" d="m149.1 308.77l198.57 40.87l8.4-39.32l-198.57-40.87Zm26.27-93.12L359.22 300L376 263.76l-183.82-84.84Zm50.95-89l156 127.78l25.74-30.52l-156-127.78ZM328 32l-33.39 23.8l120.82 160.37L448 192ZM144 400h204v-40H144Z"/>`),
		g.Group(children),
	)
}