package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24.668 18.863l-6.634 3.659l-2.792-3.729l-8.073 4.252v.94L15 19.862l2.792 3.728l7.364-4.062l.145 1.27l1.62-2.636l-.955.011v-.025l-.048.026l-2.31.027l1.059.663zM16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-1.444-14.042V8.644h-3.06v10.925l3.06-1.611zm8.653.481V8.644h-3.06v11.495l3.06-1.7zm-4.326 2.404V9.893h-3.061v8.104l2.404 3.21l.657-.364zm-8.654-.607v-8.924h-3.06v10.536l3.06-1.612z"/><path fill="currentColor" d="m18.883 20.843l-.657.364l-2.404-3.21V9.894h3.06v10.949zm-8.654-.607l-3.06 1.612V11.312h3.06v8.924z" opacity=".5"/>`),
		g.Group(children),
	)
}