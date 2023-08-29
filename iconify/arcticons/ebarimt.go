package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ebarimt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.022 36.496h8.72v-4.564A10.742 10.742 0 0 0 24 21.191h0a10.742 10.742 0 0 0-10.742 10.741v4.564h8.72"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.742 28.531h0A8.758 8.758 0 0 1 43.5 37.29v3.721h0h-17.516h0v-3.72a8.758 8.758 0 0 1 8.758-8.759Zm-21.484 0h0a8.758 8.758 0 0 1 8.758 8.758v3.721h0H4.5h0v-3.72a8.758 8.758 0 0 1 8.758-8.759Zm9.162-18.502c-.36.58 0 1.008 0 1.575c0 .925-.995 1.036-.995 2.127c0 .525.288 3.039 2.575 3.039a2.7 2.7 0 0 0 2.84-2.901c0-1.105-.856-1.312-.856-1.934c0-.787.552-.912.552-1.906"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.381 6.99c-1.09 1.616.095 2.044.095 2.934c0 .72-.476.932-.476 1.984a6.393 6.393 0 0 0 .243 1.982c.304.822 2.432 1.858 2.597-.02m-5.415-.14c0 1.783 2.708 1.534 2.708.048"/>`),
		g.Group(children),
	)
}