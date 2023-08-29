package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 0v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0 0 .96H2v1h-.52a.48.48 0 0 0 0 .96H2v1h-.52a.48.48 0 0 0 0 .96H2v2h12V0H2zm1.5 14a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zM12 6H6V3h6v3z"/>`),
		g.Group(children),
	)
}