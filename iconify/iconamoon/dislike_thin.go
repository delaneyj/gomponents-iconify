package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DislikeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 14l-.493.082A.5.5 0 0 1 15 13.5v.5ZM4 14v.5a.5.5 0 0 1-.5-.5H4Zm16.522-2.392l.49-.098l-.49.098ZM6 3.5h11.36v1H6v-1Zm12.56 11H15v-1h3.56v1Zm-3.067-.582l.806 4.835l-.986.165l-.806-4.836l.986-.164ZM14.82 20.5h-.213v-1h.213v1Zm-3.126-1.558l-2.515-3.774l.832-.555l2.515 3.774l-.832.555ZM7.93 14.5H4v-1h3.93v1ZM3.5 14V6h1v8h-1Zm16.312-8.49l1.2 6l-.98.196l-1.2-6l.98-.196ZM9.178 15.168A1.5 1.5 0 0 0 7.93 14.5v-1a2.5 2.5 0 0 1 2.08 1.113l-.832.555Zm7.121 3.585a1.5 1.5 0 0 1-1.48 1.747v-1a.5.5 0 0 0 .494-.582l.986-.165ZM18.56 13.5a1.5 1.5 0 0 0 1.471-1.794l.98-.196a2.5 2.5 0 0 1-2.45 2.99v-1Zm-1.2-10a2.5 2.5 0 0 1 2.452 2.01l-.98.196A1.5 1.5 0 0 0 17.36 4.5v-1Zm-2.754 17a3.5 3.5 0 0 1-2.913-1.558l.832-.555a2.5 2.5 0 0 0 2.08 1.113v1ZM6 4.5A1.5 1.5 0 0 0 4.5 6h-1A2.5 2.5 0 0 1 6 3.5v1Z"/><path stroke="currentColor" d="M8 14V4"/></g>`),
		g.Group(children),
	)
}