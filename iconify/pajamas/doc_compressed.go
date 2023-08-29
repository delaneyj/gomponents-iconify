package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocCompressed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 14.5V6H9.75A1.75 1.75 0 0 1 8 4.25V1.5H3.5v13H5V13h2v1.5h5.5ZM7 11h2v2H7v-2Zm0 0V9h2V7H7V5H5v2h2v2H5v2h2Zm2.5-9.379L12.379 4.5H9.75a.25.25 0 0 1-.25-.25V1.621ZM3 0a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V4.414a1 1 0 0 0-.293-.707L10.293.293A1 1 0 0 0 9.586 0H3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}