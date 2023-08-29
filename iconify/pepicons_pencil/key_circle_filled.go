package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilKeyCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M15.5 14.611V12a2.6 2.6 0 0 0-.012-.25a4.5 4.5 0 1 0-4.975 0a2.523 2.523 0 0 0-.013.25v8a2.5 2.5 0 0 0 5 0a.5.5 0 0 0-.182-.386l-.786-.646l.716-.41a.5.5 0 0 0 .252-.434v-.775a.5.5 0 0 0-.146-.353l-.998-1.001l1.003-1.036a.5.5 0 0 0 .141-.348Zm-1.047-2.986c.031.121.047.247.047.375v2.409l-1.204 1.243a.5.5 0 0 0 .005.701l1.199 1.203v.278l-1.093.625a.5.5 0 0 0-.07.82l1.147.944A1.5 1.5 0 0 1 11.5 20v-8c0-.128.016-.254.047-.375a.5.5 0 0 0-.242-.562a3.5 3.5 0 1 1 3.39 0a.5.5 0 0 0-.242.562Z"/><path d="M13 8.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilKeyCircleFilled0)"/></g>`),
		g.Group(children),
	)
}