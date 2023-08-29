package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilEyeCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M13 19c4.658 0 8.5-2.161 8.5-5S17.658 9 13 9c-4.658 0-8.5 2.161-8.5 5s3.842 5 8.5 5Zm0-9c4.179 0 7.5 1.868 7.5 4c0 2.132-3.321 4-7.5 4s-7.5-1.868-7.5-4c0-2.132 3.321-4 7.5-4Z" clip-rule="evenodd"/><path d="M12.5 6.5a.5.5 0 0 1 1 0v3a.5.5 0 0 1-1 0v-3Zm4.01.402a.5.5 0 0 1 .98.196l-.5 2.5a.5.5 0 0 1-.98-.196l.5-2.5Zm-7.02 0a.5.5 0 0 0-.98.196l.5 2.5a.5.5 0 0 0 .98-.196l-.5-2.5ZM5.429 8.243a.5.5 0 0 0-.858.514l1.5 2.5a.5.5 0 0 0 .858-.514l-1.5-2.5Zm15.142 0a.5.5 0 1 1 .858.514l-1.5 2.5a.5.5 0 1 1-.858-.514l1.5-2.5ZM16 13.5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilEyeCircleFilled0)"/></g>`),
		g.Group(children),
	)
}