package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VimNormalModeSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1.75C0 .784.784 0 1.75 0h12.5C15.216 0 16 .784 16 1.75v12.5A1.75 1.75 0 0 1 14.25 16H1.75A1.75 1.75 0 0 1 0 14.25V1.75Zm1.75-.25a.25.25 0 0 0-.25.25v12.5c0 .138.112.25.25.25h12.5a.25.25 0 0 0 .25-.25V1.75a.25.25 0 0 0-.25-.25H1.75Zm2.486 1.798a.75.75 0 0 1 .828.208l5.686 6.498V4a.75.75 0 0 1 1.5 0v8a.75.75 0 0 1-1.314.494L5.25 5.996V12a.75.75 0 0 1-1.5 0V4a.75.75 0 0 1 .486-.702Z"/>`),
		g.Group(children),
	)
}