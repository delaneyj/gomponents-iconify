package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrolleyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 1)"><circle cx="4.437" cy="12.437" r="1.437"/><ellipse cx="10.452" cy="12.485" rx="1.452" ry="1.485"/><path d="M14.684.337c-.717 0-1.297.599-1.297 1.338l-.172 1.343H.557c-.318 0-.541.287-.541.654c0 .139.01.27.072.375l1.211 5.454c.04.172.168.299.331.369a.508.508 0 0 0 .265.083H12.36l-.174 1.093H2.564c-.319 0-.501.056-.501.423s.182.474.501.474l10.79.001l1.16-8.948c.057.007.111.018.17.018c.715 0 1.295-.6 1.295-1.338c0-.74-.581-1.339-1.295-1.339zm-1.938 5.772H11V3.968h2.078l-.332 2.141zM2.17 8.994l-.309-2.035h2.182v2.035H2.17zm5.799-5.025h2.047v2.047H7.969V3.969zm-.938 2.078H4.953V3.953h2.078v2.094zm-3-.01H1.834L1.52 3.971h2.512v2.066h-.001zm.938.932h2.062v2.078H4.969V6.969zm3-.016h2.062v2.094H7.969V6.953zm3 2.078V6.984h1.906l-.289 2.047h-1.617z"/></g>`),
		g.Group(children),
	)
}