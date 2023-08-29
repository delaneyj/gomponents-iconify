package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pwa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.597 7.482L24 16.518h-2.51l-.58-1.618h-3.246l.694-1.754h2.002l-.95-2.66l1.188-3.004zm-8.111 0l1.772 5.84l2.492-5.84h2.415l-3.643 9.036H13.14l-1.64-5.237l-1.72 5.237H7.404L6.17 14.402l1.214-3.742l1.342 2.661l1.903-5.839h1.857zm-8.746 0c1.064 0 1.872.305 2.424.917a2.647 2.647 0 0 1 .28.368L5.37 12.08l-.385 1.185c-.352.1-.753.151-1.204.151H2.293v3.102H0V7.482zm-.58 1.753h-.866v2.428h.86c.557 0 .94-.12 1.148-.358c.19-.215.284-.506.284-.873c0-.364-.107-.654-.323-.871c-.216-.217-.583-.326-1.103-.326z"/>`),
		g.Group(children),
	)
}