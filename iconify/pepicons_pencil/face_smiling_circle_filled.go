package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceSmilingCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilFaceSmilingCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M12.75 20.5a7.75 7.75 0 1 0 0-15.5a7.75 7.75 0 0 0 0 15.5Zm0 1a8.75 8.75 0 1 0 0-17.5a8.75 8.75 0 0 0 0 17.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.746 14.763a.55.55 0 0 0-.994.47L9.25 15l-.498.233v.002l.001.001l.002.003l.004.01a1.353 1.353 0 0 0 .059.109a4.21 4.21 0 0 0 .895 1.06c.676.577 1.737 1.132 3.287 1.132c1.55 0 2.611-.555 3.287-1.132a4.29 4.29 0 0 0 .723-.79a3.367 3.367 0 0 0 .23-.38l.005-.009l.002-.003v-.001s.001-.002-.497-.235l.498.233a.55.55 0 0 0-.994-.47l-.004.007a2.358 2.358 0 0 1-.14.226c-.11.157-.285.37-.537.586c-.497.423-1.31.868-2.573.868s-2.076-.445-2.573-.868a3.189 3.189 0 0 1-.537-.586a2.316 2.316 0 0 1-.14-.226l-.004-.007Z" clip-rule="evenodd"/><path d="M11.5 10.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm5.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilFaceSmilingCircleFilled0)"/></g>`),
		g.Group(children),
	)
}