package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackMa0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackMa1)"><use href="#flagpackMa0"/><path fill="#C51918" fill-rule="evenodd" d="M0 0h32v22a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V0Z" clip-rule="evenodd"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#579D20" fill-rule="evenodd" d="M22.482 19.626L16.166 2.701h-.274L9.795 19.626l6.332-3.805l6.355 3.805ZM15.467 8.222l.677-2.638l.702 2.714l2.019 5.656l1.185 2.739l-2.559-1.803l-1.366-.818l-1.344.808l-2.5 1.813l1.156-2.795l2.03-5.676Z" clip-rule="evenodd"/><path fill="#579D20" fill-rule="evenodd" d="m12.662 13.473l3.496 2.324l3.263-2.324l6.195-5.237H6.386l6.276 5.237Zm.073-1.999l-2.377-1.455H21.58l-2.041 1.293l-3.397 2.577l-3.408-2.415Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackMa1"><use href="#flagpackMa0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}