package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.961.031H5.003v5.952h7.958V.031zM12 5h-2V1h2v4z"/><path d="M14.988.031h-.909v7.07H3.941V.031H2.99c-1.105 0-2 .92-2 2.052v10.895c0 1.133.896 2.053 2 2.053h12c1.106 0 2-.92 2-2.053V2.083L14.988.031zM12.968 13h-8v-1h8v1zm0-2h-8v-1h8v1z"/></g>`),
		g.Group(children),
	)
}