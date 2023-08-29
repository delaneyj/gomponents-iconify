package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Overlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 8h-4V4a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v18a2.002 2.002 0 0 0 2 2h4v4a2.002 2.002 0 0 0 2 2h18a2.002 2.002 0 0 0 2-2V10a2.002 2.002 0 0 0-2-2ZM4 22V4h18v4H10a2.002 2.002 0 0 0-2 2v12Zm18 0h-2.586L10 12.586V10h2.586L22 19.416Zm-12-6.586L16.586 22H10Zm12.001 1.173L15.414 10H22ZM10 28v-4h12a2.002 2.002 0 0 0 2-2V10h4v18Z"/>`),
		g.Group(children),
	)
}