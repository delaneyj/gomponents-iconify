package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiffModified(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M13 1H1c-.55 0-1 .45-1 1v12c0 .55.45 1 1 1h12c.55 0 1-.45 1-1V2c0-.55-.45-1-1-1zm0 13H1V2h12v12zM4 8c0-1.66 1.34-3 3-3s3 1.34 3 3s-1.34 3-3 3s-3-1.34-3-3z" fill="currentColor"/>`),
		g.Group(children),
	)
}