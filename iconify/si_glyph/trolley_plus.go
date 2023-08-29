package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrolleyPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><g fill="currentColor" transform="translate(4 12)"><circle cx="1.437" cy="1.437" r="1.437"/><ellipse cx="7.452" cy="1.485" rx="1.452" ry="1.485"/></g><path d="M6 8h2.062v2.078H6zm-.969-3.029H2.52l.314 2.066h2.197V4.971zM6 5h2.078v2.094H6zm3 0h2.047v2.047H9zm5.078-.031H12v2.14h1.746l.332-2.14zM3.17 9.994h1.873V7.959H2.861l.309 2.035zM9 8h2.062v2.094H9z"/><path fill="currentColor" d="M3.563 12.046c-.319 0-.501.056-.501.423s.182.474.501.474l9.416.001v-.897H3.563v-.001zm8.406-2.015V7.984h1.906l-.281 1.995h1.145l.775-5.984c.057.007.111.018.17.018c.715 0 1.295-.6 1.295-1.338c0-.739-.58-1.338-1.295-1.338c-.717 0-1.297.599-1.297 1.338l-.172 1.343H1.557c-.318 0-.541.287-.541.654c0 .139.01.27.072.375l1.211 5.454c.04.172.168.299.331.369a.508.508 0 0 0 .265.083H12.98v-.922h-1.011zM12 4.969h2.078l-.332 2.141H12V4.969zm-3.031 0h2.047v2.047H8.969V4.969zm-6.449.002h2.512v2.066H2.835L2.52 4.971zm2.523 5.023H3.17l-.309-2.035h2.182v2.035zm2.988.053H5.969V7.969h2.062v2.078zm0-3H5.953V4.953h2.078v2.094zm3 3H8.969V7.953h2.062v2.094zM17 12h-1.062v-1H15v1h-1v.958h1v.917h.938v-.917H17V12z"/></g>`),
		g.Group(children),
	)
}