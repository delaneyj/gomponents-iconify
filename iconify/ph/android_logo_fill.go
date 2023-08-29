package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M207.06 88.67c-.74-.74-1.49-1.46-2.24-2.17l24.84-24.84a8 8 0 0 0-11.32-11.32l-26 26a111.43 111.43 0 0 0-128.55.19L37.66 50.34a8 8 0 0 0-11.32 11.32L51.4 86.72A113.38 113.38 0 0 0 16 169.13V192a16 16 0 0 0 16 16h192a16 16 0 0 0 16-16v-24a111.25 111.25 0 0 0-32.94-79.33ZM92 168a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm72 0a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}