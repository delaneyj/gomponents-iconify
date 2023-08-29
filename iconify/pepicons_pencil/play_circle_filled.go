package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPlayCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="m9 18.321l9.014-4.883L9 7.804v10.517Zm9.49-4.003a1 1 0 0 0 .054-1.728L9.53 6.956A1 1 0 0 0 8 7.804v10.517a1 1 0 0 0 1.476.88l9.015-4.883Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPlayCircleFilled0)"/></g>`),
		g.Group(children),
	)
}