package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarGarage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M.043 2.083v11.855h1.939V3.943h12.036v9.995h1.94V2.083H.043z"/><path d="M3 5h9.956v.952H3zm0 2h9.957v.971H3zm8.97 5.96v.515c0 .271-.438.492-.975.492c-.536 0-.974-.221-.974-.492h.001v-.485H5.946v.498c0 .247-.434.449-.964.449s-.961-.202-.961-.449v-.509h.139c-1.287-.117-1.156-1.13-1.156-1.13c0-.73.152-1.901.996-2.402V8h1v1.192l.059-.008S7.096 9.001 8 9.001c.903 0 2.79.214 2.901.23l.098.018V8h1v1.457c.756.331 1 1.354 1 2.392c0 0-.027.93-1.03 1.111zm-7.001-.991h-1v-1h1v1zm3.99-.011h-2v-1h2v1zm3.01.011h-1v-1h1v1z"/></g>`),
		g.Group(children),
	)
}