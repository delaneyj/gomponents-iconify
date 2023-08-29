package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2.5.5V0H2v.5h.5Zm10 0h.5V0h-.5v.5ZM4.947 4.724a.5.5 0 0 0-.894-.448l.894.448ZM2.5 8.494l-.447-.223l-.146.293l.21.251l.383-.32Zm5 5.997l-.384.32a.5.5 0 0 0 .769 0l-.385-.32Zm5-5.996l.384.32l.21-.251l-.146-.293l-.447.224Zm-1.553-4.219a.5.5 0 0 0-.894.448l.894-.448ZM8 9.494v-.5H7v.5h1Zm-.5-4.497A4.498 4.498 0 0 1 3 .5H2a5.498 5.498 0 0 0 5.5 5.497v-1ZM2.5 1h10V0h-10v1ZM12 .5a4.498 4.498 0 0 1-4.5 4.497v1c3.038 0 5.5-2.46 5.5-5.497h-1ZM4.053 4.276l-2 3.995l.895.448l2-3.995l-.895-.448ZM2.116 8.815l5 5.996l.769-.64l-5-5.996l-.769.64Zm5.768 5.996l5-5.996l-.768-.64l-5 5.996l.769.64Zm5.064-6.54l-2-3.995l-.895.448l2 3.995l.895-.448ZM8 14.49V9.494H7v4.997h1Z"/>`),
		g.Group(children),
	)
}