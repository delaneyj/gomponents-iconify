package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whitesunbehindcloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="120.6" cy="164.349" r="98.326" fill="#FFD60D"/><path fill="#B0E9FF" d="M438.853 309.85c-5.478 0-10.796.678-15.886 1.934a68.786 68.786 0 0 0 1.756-15.399c0-37.901-30.725-68.625-68.625-68.625a68.3 68.3 0 0 0-38.756 11.995c-12.874-49.6-57.944-86.224-111.571-86.224c-63.658 0-115.263 51.605-115.263 115.263c-47.861 0-86.66 38.742-86.66 86.603S42.646 442 90.507 442h348.346c36.524 0 66.132-29.551 66.132-66.075c-.001-36.523-29.609-66.075-66.132-66.075z"/>`),
		g.Group(children),
	)
}