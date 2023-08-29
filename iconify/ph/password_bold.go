package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PasswordBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M48 56v144a12 12 0 0 1-24 0V56a12 12 0 0 1 24 0Zm82.73 50.7L116 111.48V96a12 12 0 0 0-24 0v15.48l-14.73-4.78a12 12 0 1 0-7.41 22.82l14.72 4.79l-9.1 12.52a12 12 0 1 0 19.42 14.11l9.1-12.52l9.1 12.52a12 12 0 1 0 19.42-14.11l-9.1-12.52l14.72-4.79a12 12 0 1 0-7.41-22.82Zm111.12 7.7a12 12 0 0 0-15.12-7.7L212 111.48V96a12 12 0 0 0-24 0v15.48l-14.73-4.78a12 12 0 1 0-7.41 22.82l14.72 4.79l-9.1 12.52a12 12 0 1 0 19.42 14.11l9.1-12.52l9.1 12.52a12 12 0 1 0 19.42-14.11l-9.1-12.52l14.72-4.79a12 12 0 0 0 7.71-15.12Z"/>`),
		g.Group(children),
	)
}