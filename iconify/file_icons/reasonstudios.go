package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reasonstudios(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M430.546 109.382L253.673 6.982c-16.291-9.31-34.91-9.31-51.2 0L25.6 109.382C9.31 118.69 0 134.982 0 153.6v204.8c0 18.618 9.31 34.91 25.6 44.218l176.873 102.4c16.29 9.31 34.909 9.31 51.2 0l176.872-102.4C446.836 393.31 456 377.018 456 358.4V153.6c0-18.618-9.164-34.91-25.454-44.218zM228.073 256v169.89L81.455 339.783V169.89l146.618-83.782l146.618 83.782L228.073 256z"/>`),
		g.Group(children),
	)
}