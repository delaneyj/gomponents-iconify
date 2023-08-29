package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipToStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M7 5c-.551 0-1 .449-1 1v14c0 .551.449 1 1 1h3c.551 0 1-.449 1-1v-4.875l7.094 5.625c.403.203.897.17 1.281-.063c.386-.232.625-.65.625-1.093V6.406c0-.443-.239-.86-.625-1.093a1.33 1.33 0 0 0-.688-.188c-.2 0-.408.033-.593.125L11 10.875V6c0-.551-.449-1-1-1H7z"/>`),
		g.Group(children),
	)
}