package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobywatel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.47c1.69 0 15.25-7.77 15.25-16.94v-20c-4 0-15.25-2-15.25-2s-11.26 2-15.25 2v20c0 9.17 13.56 16.94 15.25 16.94Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.23 28.204v-12.6h0a3.586 3.586 0 0 0-3.587-3.56M24.23 23.08a6.803 6.803 0 0 0 6.8-6.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.432 12.878a6.803 6.803 0 0 0-6.802 6.802m5.305-4.233h4.404m-4.693 2.416h4.405m-4.844 2.417h4.405m-6.079 2.416h4.405m-5.88 1.119l3.115 3.114m-5.655-2.096l3.115 3.115m-3.313-.018l3.965 3.964m1.419 1.419l1.324 1.324m-3.188-1.001l-1.021 1.022m3.521-2.724h1.446m-8.641-1.81l1.2 1.2m-1.2 1l1.2 1.2m-1.2 1l1.2 1.2M24.23 23.08a6.803 6.803 0 0 1-6.803-6.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.027 12.878a6.803 6.803 0 0 1 6.801 6.802m-5.304-4.233h-4.405m4.693 2.416h-4.405m4.844 2.417h-4.405m6.079 2.416H13.52m5.88 1.119l-3.115 3.114m5.655-2.096l-3.114 3.115m3.312-.018l-3.965 3.964m-1.419 1.419l-1.324 1.324m3.187-1.001l1.023 1.022m-3.521-2.724h-1.446m8.64-1.81l-1.2 1.2m1.2 1l-1.2 1.2m1.2 1l-1.2 1.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M24 10.36V8.202m-1.643 2.149l-1.08-1.87m4.366 1.87l1.08-1.87"/>`),
		g.Group(children),
	)
}