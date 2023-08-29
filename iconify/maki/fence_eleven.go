package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FenceEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.5 7H9V5h.5a.5.5 0 0 0 0-1H9V3l-.278-.555a.254.254 0 0 0-.443 0L8 3v1H7V3l-.278-.555a.254.254 0 0 0-.443 0L6 3v1H5V3l-.278-.555a.254.254 0 0 0-.443 0L4 3v1H3V3l-.278-.555a.254.254 0 0 0-.443 0L2 3v1h-.5a.5.5 0 0 0 0 1H2v2h-.5a.5.5 0 0 0 0 1H2v.5a.5.5 0 0 0 1 0V8h1v.5a.5.5 0 0 0 1 0V8h1v.5a.5.5 0 0 0 1 0V8h1v.5a.5.5 0 0 0 1 0V8h.5a.5.5 0 0 0 0-1zM3 7V5h1v2zm2 0V5h1v2zm2 0V5h1v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}