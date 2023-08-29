package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerHifiBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 20H64a20 20 0 0 0-20 20v176a20 20 0 0 0 20 20h128a20 20 0 0 0 20-20V40a20 20 0 0 0-20-20Zm-4 192H68V44h120ZM112 76a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm16 120a44 44 0 1 0-44-44a44.05 44.05 0 0 0 44 44Zm0-64a20 20 0 1 1-20 20a20 20 0 0 1 20-20Z"/>`),
		g.Group(children),
	)
}