package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerHifiFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152 160a24 24 0 1 1-24-24a24 24 0 0 1 24 24Zm56-120v176a16 16 0 0 1-16 16H64a16 16 0 0 1-16-16V40a16 16 0 0 1 16-16h128a16 16 0 0 1 16 16Zm-92 28a12 12 0 1 0 12-12a12 12 0 0 0-12 12Zm52 92a40 40 0 1 0-40 40a40 40 0 0 0 40-40Z"/>`),
		g.Group(children),
	)
}