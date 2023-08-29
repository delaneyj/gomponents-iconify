package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HushedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 57.5C16.836 59.5 4.5 47.164 4.5 32S16.836 4.5 32 4.5c15.163 0 27.5 12.336 27.5 27.5S47.163 59.5 32 59.5z"/><circle cx="20" cy="27" r="5" fill="currentColor"/><circle cx="44" cy="27" r="5" fill="currentColor"/><path fill="currentColor" d="M51.617 15.373a16.418 16.418 0 0 0-13.492-3.615c-.703.135-.193 2.27.385 2.156c4.17-.748 8.457.4 11.693 3.133c.443.387 1.955-1.205 1.414-1.674m-25.742-3.766a16.424 16.424 0 0 0-13.492 3.615c-.541.469.969 2.063 1.412 1.674a14.236 14.236 0 0 1 11.693-3.133c.578.114 1.09-2.021.387-2.156"/><circle cx="32" cy="48" r="5" fill="currentColor"/>`),
		g.Group(children),
	)
}