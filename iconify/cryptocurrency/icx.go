package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.296 22.472a8 8 0 0 0 11.176-11.176l-2.164 2.164a5 5 0 0 1-6.848 6.848zm-1.768-1.768l2.164-2.164a5 5 0 0 1 6.848-6.848l2.164-2.164A8 8 0 0 0 9.528 20.704zM16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm8-22a2 2 0 1 0 0-4a2 2 0 0 0 0 4zM8 26a2 2 0 1 0 0-4a2 2 0 0 0 0 4z"/>`),
		g.Group(children),
	)
}