package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerHifiDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 32H64a8 8 0 0 0-8 8v176a8 8 0 0 0 8 8h128a8 8 0 0 0 8-8V40a8 8 0 0 0-8-8Zm-64 152a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z" opacity=".2"/><path d="M192 24H64a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h128a16 16 0 0 0 16-16V40a16 16 0 0 0-16-16Zm0 192H64V40h128ZM116 76a12 12 0 1 1 12 12a12 12 0 0 1-12-12Zm12 116a40 40 0 1 0-40-40a40 40 0 0 0 40 40Zm0-64a24 24 0 1 1-24 24a24 24 0 0 1 24-24Z"/></g>`),
		g.Group(children),
	)
}