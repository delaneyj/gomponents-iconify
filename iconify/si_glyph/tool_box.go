package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.988 9.031V8H5.031v1.031H2.938V8H1.014v6H17V8h-1.969v1.031h-2.043zM5 5.986v.982h8v-.982h2v.982h1.987V4H1v2.968h1.974v-.982H5zm2.003-3.043H6.01V1.031h5.959v1.912h-.96v-.956H7.003v.956z"/>`),
		g.Group(children),
	)
}