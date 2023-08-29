package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464Zm5.524 6.825l4.723 4.723l2.619-2.618c1.113-1.114 1.67-1.67 1.67-2.362c0-.692-.557-1.249-1.67-2.362S14.66 6 13.968 6c-.692 0-1.248.557-2.362 1.67l-2.618 2.619Zm3.406 6.041l.257-.257l-4.724-4.724l-.257.257C6.557 12.72 6 13.276 6 13.968c0 .692.557 1.249 1.67 2.362S9.34 18 10.032 18c.692 0 1.248-.557 2.362-1.67Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}