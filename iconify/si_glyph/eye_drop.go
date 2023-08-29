package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeDrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.661 4.051c.234-1.077-.051-2.229-.86-3.046c-1.29-1.299-3.422-1.26-4.76.089C.704 2.443.663 4.589 1.954 5.889c.81.816 1.951 1.103 3.023.868l.602.586l-.58.583l.502.506l.546-.552l5.062 5.073h2.875v-2.766L8.777 5.129l.528-.533l-.5-.505l-.555.559l-.589-.599zm5.355 6.371l.005 1.623l-1.804.009l-4.545-4.71l1.547-1.625l4.797 4.703zm2.969 4.418c0 .619-.44 1.122-.986 1.122c-.545 0-.985-.503-.985-1.122c0-.62.985-1.777.985-1.777s.986 1.157.986 1.777z"/>`),
		g.Group(children),
	)
}