package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CertificateCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 16h6v2H6zm0-4h10v2H6zm0-4h10v2H6z"/><path fill="currentColor" d="M14 26H4V6h24v10h2V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h10Z"/><path fill="currentColor" d="M22 25.59L19.41 23L18 24.41l4 4l8-8L28.59 19L22 25.59z"/>`),
		g.Group(children),
	)
}