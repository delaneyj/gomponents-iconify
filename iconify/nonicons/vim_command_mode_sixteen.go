package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VimCommandModeSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1.75C0 .784.784 0 1.75 0h12.5C15.216 0 16 .784 16 1.75v12.5A1.75 1.75 0 0 1 14.25 16H1.75A1.75 1.75 0 0 1 0 14.25V1.75Zm1.75-.25a.25.25 0 0 0-.25.25v12.5c0 .138.112.25.25.25h12.5a.25.25 0 0 0 .25-.25V1.75a.25.25 0 0 0-.25-.25H1.75Zm6.5 3.25A3.25 3.25 0 1 0 10.812 10a.75.75 0 1 1 1.182.924a4.75 4.75 0 1 1 0-5.847a.75.75 0 0 1-1.182.923A3.243 3.243 0 0 0 8.25 4.75Z"/>`),
		g.Group(children),
	)
}