package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M12.688 5.125c-.202 0-.41.033-.594.125L3.78 11.844c-.4.376-.719.671-.719 1.156c0 .485.271.76.72 1.156l8.312 6.594c.403.203.897.17 1.281-.063c.386-.232.625-.65.625-1.093v-3.688l6.094 4.844c.403.203.897.17 1.281-.063c.386-.232.625-.65.625-1.093V6.406c0-.443-.239-.86-.625-1.093a1.33 1.33 0 0 0-.688-.188c-.2 0-.408.033-.593.125L14 10.094V6.406c0-.443-.239-.86-.625-1.093a1.33 1.33 0 0 0-.688-.188z"/>`),
		g.Group(children),
	)
}