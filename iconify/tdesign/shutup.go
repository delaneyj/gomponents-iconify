package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shutup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm5-3h4v2H6V9Zm8 0h4v2h-4V9Zm-3 3.586l1 1l1-1L14.414 14l-1 1l1 1L13 17.414l-1-1l-1 1L9.586 16l1-1l-1-1L11 12.586Z"/>`),
		g.Group(children),
	)
}