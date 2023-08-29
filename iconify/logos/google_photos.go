package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePhotos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#FBBB05" d="M64 58.149c35.328 0 64 28.672 64 64V128H5.851C2.633 128 0 125.367 0 122.149c0-35.328 28.672-64 64-64Z"/><path fill="#E94335" d="M197.851 64c0 35.328-28.672 64-64 64H128V5.851C128 2.633 130.633 0 133.851 0c35.328 0 64 28.672 64 64Z"/><path fill="#4285F4" d="M192 197.851c-35.328 0-64-28.672-64-64V128h122.149c3.218 0 5.851 2.633 5.851 5.851c0 35.328-28.672 64-64 64Z"/><path fill="#0F9D58" d="M58.149 192c0-35.328 28.672-64 64-64H128v122.149c0 3.218-2.633 5.851-5.851 5.851c-35.328 0-64-28.672-64-64Z"/>`),
		g.Group(children),
	)
}