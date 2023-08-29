package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFlagSyria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#EEE" d="M0 13h36v10H0z"/><path fill="#CE1126" d="M32 5H4a4 4 0 0 0-4 4v4h36V9a4 4 0 0 0-4-4z"/><path fill="#141414" d="M32 31H4a4 4 0 0 1-4-4v-4h36v4a4 4 0 0 1-4 4z"/><path d="M9.177 18.332l-.633 1.947l1.656-1.203l1.656 1.203l-.633-1.947l1.656-1.202h-2.047l-.632-1.947l-.632 1.947H7.521zm15.6 0l-.633 1.947l1.656-1.203l1.656 1.203l-.633-1.947l1.656-1.202h-2.047l-.632-1.947l-.632 1.947h-2.047z" fill="#007A3D"/>`),
		g.Group(children),
	)
}