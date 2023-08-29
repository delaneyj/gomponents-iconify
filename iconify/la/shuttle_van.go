package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuttleVan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5C3.355 5 2 6.355 2 8v17h3.156c.446 1.719 1.992 3 3.844 3c1.852 0 3.398-1.281 3.844-3h7.312c.446 1.719 1.992 3 3.844 3c1.852 0 3.398-1.281 3.844-3H31V14.562a3.03 3.03 0 0 0-.469-1.593l-4.093-6.563A2.982 2.982 0 0 0 23.905 5zm0 2h5v6H4V8c0-.563.438-1 1-1zm7 0h6v6h-6zm8 0h3.906c.348 0 .66.176.844.469L28.188 13H20zM4 15h25v2h-3v2h3v4h-1.156c-.446-1.719-1.992-3-3.844-3c-1.852 0-3.398 1.281-3.844 3h-7.312c-.446-1.719-1.992-3-3.844-3c-1.852 0-3.398 1.281-3.844 3H4zm5 7c1.117 0 2 .883 2 2s-.883 2-2 2s-2-.883-2-2s.883-2 2-2zm15 0c1.117 0 2 .883 2 2s-.883 2-2 2s-2-.883-2-2s.883-2 2-2z"/>`),
		g.Group(children),
	)
}