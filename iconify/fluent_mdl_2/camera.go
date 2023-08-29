package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1536H0V256h549l128-128h694l128 128h549zm-128 128h-475l-128-128H731L603 384H128v1280h1792V384zm-896 128q106 0 199 40t163 109t110 163t40 200q0 106-40 199t-109 163t-163 110t-200 40q-106 0-199-40t-163-109t-110-163t-40-200q0-106 40-199t109-163t163-110t200-40zm0 896q79 0 149-30t122-83t82-122t31-149q0-79-30-149t-83-122t-122-82t-149-31q-79 0-149 30t-122 83t-82 122t-31 149q0 79 30 149t83 122t122 82t149 31zM320 512q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45q0-26 19-45t45-19z"/>`),
		g.Group(children),
	)
}