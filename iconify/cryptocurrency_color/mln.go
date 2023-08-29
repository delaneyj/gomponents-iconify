package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mln(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#0B1529"/><path fill="#FFF" d="m8.627 20.124l5.272 3.092v1.729L7 20.885v-10.59L16 5l9 5.295v10.59l-6.899 4.06v-1.73l5.271-3.091l-1.36-.786l.178-.309l1.343.776v-7.852l-6.82 4.042v11.56L16 28l-.714-.445v-11.56l-6.819-4.042v7.851l1.342-.775l.178.31l-1.36.785zm7.175-13.31L9.18 10.699L16 14.742l6.82-4.043l-6.661-3.908v1.604h-.357V6.814zm0 2.592h.357v1.657h-.357V9.406zm0 2.668h.357v1.657l-.179.081l-.178-.08v-1.658zm5.56 6.45l-.178.31l-1.436-.83l.179-.309l1.435.83zm-2.265-1.334l-.179.31l-1.435-.83l.02-.194l.158-.114l1.436.828zm-8.46 1.334l1.435-.829l.179.31l-1.435.828l-.179-.309zm2.266-1.334l1.435-.828l.16.114l.018.195l-1.435.828l-.178-.309z"/></g>`),
		g.Group(children),
	)
}