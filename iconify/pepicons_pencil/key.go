package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12.5 11.611V9a2.6 2.6 0 0 0-.012-.25a4.5 4.5 0 1 0-4.975 0A2.523 2.523 0 0 0 7.5 9v8a2.5 2.5 0 0 0 5 0a.5.5 0 0 0-.182-.386l-.786-.646l.716-.41a.5.5 0 0 0 .252-.434v-.775a.5.5 0 0 0-.146-.353l-.998-1.001l1.003-1.036a.5.5 0 0 0 .141-.348Zm-1.047-2.986c.031.121.047.247.047.375v2.409l-1.204 1.243a.5.5 0 0 0 .005.701l1.199 1.203v.278l-1.093.626a.5.5 0 0 0-.07.82l1.147.943A1.5 1.5 0 0 1 8.5 17V9c0-.128.016-.254.047-.375a.5.5 0 0 0-.242-.562a3.5 3.5 0 1 1 3.39 0a.5.5 0 0 0-.242.562Z"/><path d="M10 5.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Z"/></g>`),
		g.Group(children),
	)
}