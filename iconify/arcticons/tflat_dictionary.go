package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TflatDictionary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M31.682 30.294c.478-1.784 4.583-5.328 7.707-6.166l-3.854-6.422l-2.826 1.284"/><path d="M34.428 28.572c1.456-1.456 5.935-3.041 8.42-3.707l-5.514-7.159m0 0l-1.336.77m.778 9.976c1.876-1.083 5.289-1.805 6.724-1.754m0 0l-1.568-1.568m-10.25 5.164l9.763-1.027m0 0l-1.227-2.097M4.5 22.073h5.447m-2.724 8.221v-8.221m4.676 4.111h2.672m-2.672 4.11v-8.221h4.111"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.962 22.073v7.194c0 .567.46 1.028 1.028 1.028h.308m6.064-2.055c0 1.134-.92 2.054-2.056 2.054h0a2.055 2.055 0 0 1-2.055-2.055v-1.336c0-1.135.92-2.055 2.055-2.055h0c1.135 0 2.056.92 2.056 2.055m0 3.391v-5.446m3.031-1.696v6.115c0 .567.46 1.028 1.028 1.028h.308m-2.415-5.447h2.158"/>`),
		g.Group(children),
	)
}