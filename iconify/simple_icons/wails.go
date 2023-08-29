package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wails(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.67 5.252l-7.856 5.039l-.369-.459l-8.69-.283l1.891 1.904l5.221.106l1.63 1.656l-5.878.662l1.77 1.783l5.34-1.185l.003-.006l.993 1.168l-3.079 3.11l7.399.001l-1.582-5.002l2.209 3.14H24l-5.385-2.415h4.121l-5.384-2.418h4.117L16.297 9.73l1.088-1.443l4.09-1.076l-3.467.248l1.662-2.207zm-3.635 2.322l-6.039.43l1.455 1.826l1.813-.476l2.771-1.78zm-.252 2.84l-.86 1.145l-.001-.002l.154-.205l.707-.938zM0 12.2l5.615 1.033l-1.017-1.027L0 12.2z"/>`),
		g.Group(children),
	)
}