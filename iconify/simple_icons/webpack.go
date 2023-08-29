package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webpack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.199 18.498L12.429 24v-4.285l6.087-3.334l3.683 2.117zm.668-.603V6.388l-3.575 2.055v7.396zm-21.066.603L11.571 24v-4.285L5.484 16.38l-3.683 2.117zm-.668-.603V6.388l3.575 2.055v7.396zm.418-12.251L11.571 0v4.143L5.152 7.66l-.049.028zm20.898 0L12.429 0v4.143l6.419 3.516l.049.028l3.552-2.043zM11.57 18.74l-6.005-3.287V8.938l6.005 3.453v6.349zm.858 0l6.005-3.287V8.938l-6.005 3.453zM5.972 8.185l6.03-3.302l6.028 3.301l-6.029 3.467z"/>`),
		g.Group(children),
	)
}