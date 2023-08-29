package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSettingsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.414 5H21a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h7.414l2 2Zm-3.823 8.809l-.991.572l1 1.731l.991-.572c.393.371.872.653 1.405.811v1.145h1.998V16.35a3.495 3.495 0 0 0 1.405-.811l.992.572l.999-1.73l-.991-.573a3.507 3.507 0 0 0 0-1.622l.991-.572l-1-1.732l-.992.573a3.495 3.495 0 0 0-1.404-.812V8.5h-1.999v1.144a3.494 3.494 0 0 0-1.404.812L8.6 9.883l-1 1.732l.991.572a3.508 3.508 0 0 0 0 1.622Zm3.404.688a1.5 1.5 0 1 1 0-2.998a1.5 1.5 0 0 1 0 2.998Z"/>`),
		g.Group(children),
	)
}