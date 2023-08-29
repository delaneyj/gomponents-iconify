package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EslintOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m.5 7.5l-.447-.224a.5.5 0 0 0 0 .448L.5 7.5Zm3-6V1a.5.5 0 0 0-.447.276L3.5 1.5Zm8 0l.447-.224A.5.5 0 0 0 11.5 1v.5Zm3 6l.447.224a.5.5 0 0 0 0-.448L14.5 7.5Zm-3 6v.5a.5.5 0 0 0 .447-.276L11.5 13.5Zm-8 0l-.447.224A.5.5 0 0 0 3.5 14v-.5Zm4-9.5l.277-.416l-.277-.185l-.277.185L7.5 4Zm-3 2l-.277-.416L4 5.732V6h.5Zm0 3H4v.268l.223.148L4.5 9Zm3 2l-.277.416l.277.185l.277-.185L7.5 11Zm3-2l.277.416l.223-.148V9h-.5Zm0-3h.5v-.268l-.223-.148L10.5 6ZM.947 7.724l3-6l-.894-.448l-3 6l.894.448ZM3.5 2h8V1h-8v1Zm7.553-.276l3 6l.894-.448l-3-6l-.894.448Zm3 5.552l-3 6l.894.448l3-6l-.894-.448ZM11.5 13h-8v1h8v-1Zm-7.553.276l-3-6l-.894.448l3 6l.894-.448Zm3.276-9.692l-3 2l.554.832l3-2l-.554-.832ZM4 6v3h1V6H4Zm.223 3.416l3 2l.554-.832l-3-2l-.554.832Zm3.554 2l3-2l-.554-.832l-3 2l.554.832ZM11 9V6h-1v3h1Zm-.223-3.416l-3-2l-.554.832l3 2l.554-.832Z"/>`),
		g.Group(children),
	)
}