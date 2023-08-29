package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragDrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.999 15V2h13v5h-2V4h-9v9h3v2h-5Zm6 5V8h12v5.5h-2V10h-8v8h3.5v2h-5.5Zm8.778 3.684L13.41 13.378l10.258 3.407l-4.656 2.227l-2.235 4.672Z"/>`),
		g.Group(children),
	)
}