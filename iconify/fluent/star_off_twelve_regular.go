package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOffTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m1.854 1.146l9 9a.5.5 0 0 1-.707.708l-.772-.772a.802.802 0 0 1-1.15.515L6 9.427l-2.224 1.17a.8.8 0 0 1-1.16-.844l.424-2.475l-1.799-1.754a.8.8 0 0 1 .444-1.364l1.542-.225l-2.081-2.081a.5.5 0 1 1 .707-.708ZM8.23 8.938l-4.12-4.12l-1.88.273l1.584 1.544a.8.8 0 0 1 .23.708l-.374 2.18l1.958-1.03a.8.8 0 0 1 .744 0l1.958 1.03l-.1-.585ZM9.77 5.09L8.475 6.354l.707.707l1.577-1.537a.8.8 0 0 0-.443-1.364L7.83 3.798L6.718 1.546a.8.8 0 0 0-1.435 0L4.75 2.628l.747.747l.505-1.023l.979 1.983a.8.8 0 0 0 .602.438l2.19.318Z"/>`),
		g.Group(children),
	)
}