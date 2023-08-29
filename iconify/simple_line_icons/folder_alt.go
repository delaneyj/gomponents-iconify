package simple_line_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.16 208h-480l-96-96h-320c-35.344 0-64 28.656-64 64v160h-.304v64H.16v448c0 35.344 28.656 64 64 64h896c35.344 0 64-28.656 64-64V272c0-35.344-28.656-64-64-64zM64.145 176h290.75l78.624 77.248L453.632 272H960.16v64h-896V176h-.016zm-.001 672V400h896v448h-896z"/>`),
		g.Group(children),
	)
}