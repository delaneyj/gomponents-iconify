package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearbyOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.8 16L17 14.2l2.2-2.2L12 4.8L9.8 7L8 5.2l2.55-2.55q.3-.3.675-.463T12 2.026q.4 0 .763.15t.662.45L21.4 10.6q.6.575.6 1.388t-.6 1.412L18.8 16Zm1.025 6.575L16.05 18.8l-2.65 2.6q-.3.275-.663.425t-.737.15q-.375 0-.738-.138T10.6 21.4l-7.975-7.975q-.3-.3-.45-.663t-.15-.737q0-.375.15-.738t.45-.662l2.575-2.6L1.4 4.2l1.4-1.4l18.4 18.4l-1.375 1.375Zm-15-10.55L12 19.2l2.25-2.2l-1.425-1.425L12 16.4l-4.375-4.375l.825-.825l-1.4-1.4l-2.225 2.225ZM15.6 12.8l-4.4-4.4l.8-.8l4.4 4.4l-.8.8Z"/>`),
		g.Group(children),
	)
}