package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 11.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm-1.864 1.602a.75.75 0 0 0-1.272 0l-2.5 4A.75.75 0 0 0 8 18.25h8a.75.75 0 0 0 .6-1.2l-1.5-2a.75.75 0 0 0-1.2 0l-.844 1.125l-1.92-3.073Z"/><path fill="currentColor" fill-rule="evenodd" d="M7 2.25A2.75 2.75 0 0 0 4.25 5v14A2.75 2.75 0 0 0 7 21.75h10A2.75 2.75 0 0 0 19.75 19V7.968c0-.381-.124-.751-.354-1.055l-2.998-3.968a1.75 1.75 0 0 0-1.396-.695H7ZM5.75 5c0-.69.56-1.25 1.25-1.25h7.25v4.397c0 .414.336.75.75.75h3.25V19c0 .69-.56 1.25-1.25 1.25H7c-.69 0-1.25-.56-1.25-1.25V5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}