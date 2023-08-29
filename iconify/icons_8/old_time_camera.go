package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OldTimeCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m11.5 6l-.313.406L10 8H9V7H5v1H3v18h26V8h-7l-1.188-1.594L20.5 6h-9zm1 2h7l1.188 1.594L21 10h6v4h-5.813c-1.042-1.784-2.98-3-5.187-3c-2.206 0-4.145 1.216-5.188 3H5v-4h6l.313-.406L12.5 8zM23 11v2h2v-2h-2zm-7 2c2.22 0 4 1.78 4 4c0 2.22-1.78 4-4 4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4zM5 16h5.094A6.106 6.106 0 0 0 10 17c0 3.302 2.698 6 6 6s6-2.698 6-6c0-.337-.04-.678-.094-1H27v8H5v-8z"/>`),
		g.Group(children),
	)
}