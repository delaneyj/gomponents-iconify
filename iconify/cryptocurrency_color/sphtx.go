package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sphtx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#00b098"/><path fill="#fff" fill-rule="nonzero" d="M6 12.391v-.356h8.417v.356zm0-1.035V11h8.417v.356zm3.55 10.64v-8.893h.338v8.893zm.98 0v-8.893h.337v8.893zm15.232-.728l-3.798-4.013l.238-.251L26 21.016zM16.519 11l3.798 4.013l-.238.251l-3.798-4.012zm8.55 11l-3.798-4.013l.238-.252l3.798 4.013zm-9.242-10.268l3.798 4.012l-.239.252l-3.798-4.012zm5.421 4.768l.239-.252l4.275-4.516l.238.252l-4.275 4.516l-.238.252l-.454.48l-.239.252L16.52 22l-.238-.252l4.275-4.516l.238-.252zm-.454-.984L25.07 11l.238.252l-4.275 4.516l-.238.252l-.454.48l-.239.252l-4.274 4.516l-.239-.252l4.275-4.516l.238-.252l.455-.48z"/></g>`),
		g.Group(children),
	)
}