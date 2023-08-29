package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletinNoticeCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilBulletinNoticeCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M21 9.923H5v11h16v-11Zm-16-1a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1H5Z"/><path d="M9 12.423a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm-1 3a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Zm1 3a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm5.768-13.025a2.5 2.5 0 0 0-3.536 0L7.354 9.277a.5.5 0 1 1-.708-.707l3.88-3.88a3.5 3.5 0 0 1 4.949 0l3.879 3.88a.5.5 0 1 1-.708.707l-3.878-3.88Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilBulletinNoticeCircleFilled0)"/></g>`),
		g.Group(children),
	)
}