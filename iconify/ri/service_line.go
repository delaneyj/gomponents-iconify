package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServiceLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.161 4.47a6.5 6.5 0 0 1 8.84-.329a6.5 6.5 0 0 1 9.178 9.154l-7.765 7.791a2 2 0 0 1-2.719.102l-.11-.102l-7.764-7.79a6.5 6.5 0 0 1 .34-8.827Zm1.414 1.413a4.5 4.5 0 0 0-.146 6.211l.146.153L12 19.672l5.303-5.304l-3.535-3.535l-1.06 1.06a3 3 0 0 1-4.244-4.242l2.102-2.102a4.501 4.501 0 0 0-5.837.188l-.154.146Zm8.486 2.829a1 1 0 0 1 1.414 0l4.242 4.242l.708-.707a4.5 4.5 0 0 0-6.211-6.51l-.153.146l-3.182 3.182a1 1 0 0 0-.078 1.327l.078.088a1 1 0 0 0 1.327.077l.087-.077l1.768-1.768Z"/>`),
		g.Group(children),
	)
}