package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmSlateThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 108H71.26L209 71.64a4 4 0 0 0 2.84-4.92l-8.16-30a11.92 11.92 0 0 0-14.54-8.31L36.83 68.61a11.78 11.78 0 0 0-7.26 5.53A11.59 11.59 0 0 0 28.41 83l7.67 28.19a3.87 3.87 0 0 0-.08.81v88a12 12 0 0 0 12 12h160a12 12 0 0 0 12-12v-88a4 4 0 0 0-4-4Zm-90-54.67l38 21.92l-47.92 12.65L78.12 66Zm65.16-17.2a4 4 0 0 1 1-.13a3.88 3.88 0 0 1 3.84 2.81l7.1 26.12l-28.1 7.42l-38-21.92Zm-154.7 42a3.86 3.86 0 0 1 2.37-1.8l28.27-7.46l38 21.91l-61.85 16.34l-7.12-26.18a3.66 3.66 0 0 1 .37-2.8ZM212 200a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4v-84h168Z"/>`),
		g.Group(children),
	)
}