package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PubgMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.014 14.304A16.776 16.776 0 0 1 43.5 25.271"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 25.27a53.822 53.822 0 0 1-1.136 10.553m-28.35-21.519l-.858 1.253a2.792 2.792 0 0 0-2.33 2.19"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.155 15.557a36.262 36.262 0 0 1 12.956 2.363c6.779 2.602 9.247 4.799 10.129 6.083a2.208 2.208 0 0 1-.313 2.919a2.097 2.097 0 0 1-2.729.115"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.962 24.742c-4.904-4.675-21.623-8.117-22.478-6.71c-.529.87-.963 2.107-.615 2.57"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.869 20.602s-1.89 3.107-1.248 3.56c0 0-.614-.069-.92 1.246s-.307 1.847.52 2.416s1.794 1.4 1.794 1.4l-2.043 5.084s-.246 2.108 4.457 3.292a3.796 3.796 0 0 0 4.543-1.822m7.414 2.8c2.741 1.33 6.19 1.16 7.182.096a13.815 13.815 0 0 0 1.909-3.25s1.195 2.412 5.15 2.373a13.051 13.051 0 0 0 6.737-1.974"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.772 32.21s-7.694-4.17-13.208-5.033a46.129 46.129 0 0 0-9.064-.64m1.12-2.375a24.34 24.34 0 0 1 6.869.26a35.3 35.3 0 0 1 9.835 3.155l1.357-3.137a42.687 42.687 0 0 0-11.127-3.515a32.769 32.769 0 0 0-5.685-.324m23.608 14.823l2.72-8.387"/><circle cx="26.149" cy="25.411" r=".75" fill="currentColor"/><circle cx="24.502" cy="28.793" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.386 38.578s-12.163-5.699-16.414-4.27"/>`),
		g.Group(children),
	)
}