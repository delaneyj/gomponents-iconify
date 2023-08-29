package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.253 9.542c-.388-.416-2.265-2.271-3.122-3.118A1.49 1.49 0 0 0 15.098 6H4.902c-.394 0-.77.165-1.033.424c-.858.847-2.734 2.701-3.122 3.118c-.485.521-.723.902-.624 1.449s.466 2.654.556 3.074c.088.419.684.935 1.24.935h16.162c.556 0 1.152-.516 1.241-.935c.089-.42.457-2.527.556-3.074s-.139-.929-.625-1.449zm-5.239.461a.271.271 0 0 0-.238.133L12.966 12H7.034l-.809-1.864a.271.271 0 0 0-.238-.133H2.473L4.495 8h11.01l2.023 2.003h-3.514z"/>`),
		g.Group(children),
	)
}