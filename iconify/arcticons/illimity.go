package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Illimity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.841 9.837C15.686-1.558-5.587 21.34 10.183 33.194c2.032 1.528 5.196 2.407 7.583 3.479c1.958.88 3.132 2.04 4.955 2.506c14.078 3.596 26.542-3.55 18.24-22.343c-.743-1.682-3.383-4.555-5.018-5.397c-13.219-6.805-10.338 12.107-11.932 18.424c-.649 2.572-1.489 3.955-2.909 5.271c-1.033.957-2.911 2.409-3.669 1.222c-2.452-3.841-5.449-11.102.465-13.407c3.312-1.291 7.963 2.545 11.721 1.054c7.688-3.049 5.434-8.915 3.162-13.787"/>`),
		g.Group(children),
	)
}