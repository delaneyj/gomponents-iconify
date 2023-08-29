package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileHtmlBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 152a12 12 0 0 1-12 12h-2v44a12 12 0 0 1-24 0v-44h-2a12 12 0 0 1 0-24h28a12 12 0 0 1 12 12Zm-72-12a12 12 0 0 0-12 12v16h-8v-16a12 12 0 0 0-24 0v56a12 12 0 0 0 24 0v-16h8v16a12 12 0 0 0 24 0v-56a12 12 0 0 0-12-12Zm188 56h-4v-44a12 12 0 0 0-24 0v56a12 12 0 0 0 12 12h16a12 12 0 0 0 0-24Zm-44.8-55.56a12 12 0 0 0-13.49 5.39L172 168.68l-13.71-22.85A12 12 0 0 0 136 152v56a12 12 0 0 0 24 0v-12.68l1.71 2.85a12 12 0 0 0 20.58 0l1.71-2.85V208a12 12 0 0 0 24 0v-56a12 12 0 0 0-8.8-11.56ZM208 120a12 12 0 0 1-12-12v-4h-48a12 12 0 0 1-12-12V44H60v64a12 12 0 0 1-24 0V40a20 20 0 0 1 20-20h96a12 12 0 0 1 8.49 3.52l56 56A12 12 0 0 1 220 88v20a12 12 0 0 1-12 12Zm-25-40l-23-23v23Z"/>`),
		g.Group(children),
	)
}