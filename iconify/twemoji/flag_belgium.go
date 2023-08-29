package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBelgium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#141414" d="M7 5a4 4 0 0 0-4 4v18a4 4 0 0 0 4 4h6V5H7z"/><path fill="#FDDA24" d="M13 5h10v26H13z"/><path fill="#EF3340" d="M29 5h-6v26h6a4 4 0 0 0 4-4V9a4 4 0 0 0-4-4z"/>`),
		g.Group(children),
	)
}