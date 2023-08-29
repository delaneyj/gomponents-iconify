package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 5v20h5v5.094l1.625-1.313L13.344 25H30V5zm2 2h24v16H12.656l-.281.219L9 25.906V23H4zm4 5v2h16v-2zm0 4v2h12v-2z"/>`),
		g.Group(children),
	)
}