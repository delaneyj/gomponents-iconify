package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmReelDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M128 32a96 96 0 1 0 96 96a96 96 0 0 0-96-96ZM80 144a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm48 48a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm0-96a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm48 48a16 16 0 1 1 16-16a16 16 0 0 1-16 16Z" opacity=".2"/><path d="M224 216h-40.64A103.95 103.95 0 1 0 128 232h96a8 8 0 0 0 0-16ZM40 128a88 88 0 1 1 88 88a88.1 88.1 0 0 1-88-88Zm88-24a24 24 0 1 0-24-24a24 24 0 0 0 24 24Zm0-32a8 8 0 1 1-8 8a8 8 0 0 1 8-8Zm24 104a24 24 0 1 0-24 24a24 24 0 0 0 24-24Zm-32 0a8 8 0 1 1 8 8a8 8 0 0 1-8-8Zm56-24a24 24 0 1 0-24-24a24 24 0 0 0 24 24Zm0-32a8 8 0 1 1-8 8a8 8 0 0 1 8-8Zm-96-16a24 24 0 1 0 24 24a24 24 0 0 0-24-24Zm0 32a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/></g>`),
		g.Group(children),
	)
}