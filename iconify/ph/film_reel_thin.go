package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmReelThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 220h-56.82a100 100 0 1 0-39.18 8h96a4 4 0 0 0 0-8ZM36 128a92 92 0 1 1 92 92a92.1 92.1 0 0 1-92-92Zm92-28a20 20 0 1 0-20-20a20 20 0 0 0 20 20Zm0-32a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm20 108a20 20 0 1 0-20 20a20 20 0 0 0 20-20Zm-32 0a12 12 0 1 1 12 12a12 12 0 0 1-12-12Zm60-28a20 20 0 1 0-20-20a20 20 0 0 0 20 20Zm0-32a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm-96-8a20 20 0 1 0 20 20a20 20 0 0 0-20-20Zm0 32a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}