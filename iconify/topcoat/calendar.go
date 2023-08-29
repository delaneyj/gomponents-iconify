package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M37.499 5.5H35.5v5.52c0 1.23-.234 1.48-1.484 1.48h-6.017c-1.25 0-1.499-.271-1.499-1.5V5.5h-11v5.52c0 1.23-.25 1.48-1.5 1.48H8c-1.25 0-1.5-.271-1.5-1.5V5.5h-2c-2.58 0-3 .561-3 3v28.998c-.01 2.43.6 3.002 3 3.002h32.999c2.55 0 3.001-.531 3.001-3.002V8.5c0-2.35-.381-3-3.001-3zm-1.999 30h-29v-20h29v20z"/><path fill="currentColor" d="M10 10.5h2c1.25 0 1.5-.25 1.5-1.48v-5c0-1.25-.27-1.5-1.5-1.5l-2-.02c-1.189 0-1.5.28-1.5 1.5v5c0 1.229.25 1.5 1.5 1.5zm19.999 0h2c1.25 0 1.501-.25 1.501-1.48v-5c0-1.25-.27-1.5-1.5-1.5l-2-.02c-1.189 0-1.5.28-1.5 1.5v5c0 1.229.249 1.5 1.499 1.5zm-19.499 9h5v5h-5zm8 0h5v5h-5zm8 0h5v5h-5zm-16 8h5v5h-5zm8 0h5v5h-5zm8 0h5v5h-5z"/>`),
		g.Group(children),
	)
}