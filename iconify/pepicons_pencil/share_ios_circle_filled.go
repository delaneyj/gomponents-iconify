package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareIosCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilShareIosCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M12.998 5.5a.5.5 0 1 1 1 0L14 16a.5.5 0 1 1-1 0l-.002-10.5Z"/><path d="M13.82 5.884a.5.5 0 0 0-.64-.768l-3 2.5a.5.5 0 1 0 .64.768l3-2.5Z"/><path d="M13.18 5.884a.5.5 0 1 1 .64-.768l3 2.5a.5.5 0 0 1-.64.768l-3-2.5ZM16.248 12a.5.5 0 0 1 0-1h.917c1.271 0 2.333.891 2.333 2.039v6.923c0 1.147-1.062 2.038-2.333 2.038H9.832c-1.271 0-2.334-.891-2.334-2.038v-6.923C7.498 11.89 8.561 11 9.832 11h.916a.5.5 0 0 1 0 1h-.916c-.754 0-1.334.486-1.334 1.039v6.923c0 .552.58 1.038 1.334 1.038h7.333c.754 0 1.333-.486 1.333-1.038v-6.923c0-.553-.58-1.039-1.333-1.039h-.917Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilShareIosCircleFilled0)"/></g>`),
		g.Group(children),
	)
}