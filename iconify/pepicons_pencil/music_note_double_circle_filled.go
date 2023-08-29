package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteDoubleCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilMusicNoteDoubleCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="m18.139 4.376l-7 .52A1.5 1.5 0 0 0 9.75 6.392v.737a2.324 2.324 0 0 0 .005.117a1.5 1.5 0 0 0 1.612 1.378l7-.55A1.5 1.5 0 0 0 19.75 6.58v-.707c0-.056 0-.056-.004-.112a1.5 1.5 0 0 0-1.607-1.384ZM10.75 7.129v-.737a.5.5 0 0 1 .463-.498l7-.521a.5.5 0 0 1 .537.499v.707a.5.5 0 0 1-.46.499l-7 .55a.5.5 0 0 1-.538-.46l-.002-.04ZM4.75 19c0 1.408 1.365 2.5 3 2.5s3-1.092 3-2.5s-1.365-2.5-3-2.5s-3 1.092-3 2.5Zm5 0c0 .8-.874 1.5-2 1.5s-2-.7-2-1.5s.874-1.5 2-1.5s2 .7 2 1.5Zm4.92.805c.556.445 1.296.695 2.08.695c1.635 0 3-1.092 3-2.5s-1.365-2.5-3-2.5s-3 1.092-3 2.5c0 .697.34 1.341.92 1.805ZM18.75 18c0 .8-.874 1.5-2 1.5c-.56 0-1.08-.176-1.455-.476c-.353-.282-.545-.646-.545-1.024c0-.8.874-1.5 2-1.5s2 .7 2 1.5Z" clip-rule="evenodd"/><path d="M9.75 6.5a.5.5 0 0 1 1 0V19a.5.5 0 0 1-1 0V6.5Zm9 0a.5.5 0 0 1 1 0V18a.5.5 0 0 1-1 0V6.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilMusicNoteDoubleCircleFilled0)"/></g>`),
		g.Group(children),
	)
}