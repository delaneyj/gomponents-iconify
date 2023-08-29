package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderArrowLeftRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 19c0 .34.04.67.09 1H4a2 2 0 0 1-2-2V6c0-1.11.89-2 2-2h6l2 2h8a2 2 0 0 1 2 2v5.81c-.61-.35-1.28-.59-2-.72V8H4v10h9.09c-.05.33-.09.66-.09 1m10-2l-3-2.5V16h-4v2h4v1.5l3-2.5m-5 1.5L15 21l3 2.5V22h4v-2h-4v-1.5Z"/>`),
		g.Group(children),
	)
}