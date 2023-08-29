package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudPakMulticloudMgmt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 19a1.982 1.982 0 0 0-.512.074l-1.781-1.781l-.026.026a2.888 2.888 0 0 0 0-2.638l.026.026l1.781-1.781a2.034 2.034 0 1 0-1.414-1.414l-1.781 1.781l.026.026a2.888 2.888 0 0 0-2.638 0l.026-.026l-1.781-1.781a2.034 2.034 0 1 0-1.414 1.414l1.781 1.781l.026-.026a2.887 2.887 0 0 0 0 2.638l-.026-.026l-1.781 1.781a2.034 2.034 0 1 0 1.414 1.414l1.781-1.781l-.026-.026a2.887 2.887 0 0 0 2.638 0l-.026.026l1.781 1.781A1.996 1.996 0 1 0 21 19Zm-5-2a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M16 31a.999.999 0 0 1-.504-.136l-12-7A1 1 0 0 1 3 23V9a1 1 0 0 1 .496-.864l12-7a1 1 0 0 1 1.008 0l12 7l-1.008 1.728L16 3.158L5 9.574v12.852l11 6.417l11-6.417V15h2v8a1 1 0 0 1-.496.864l-12 7A.999.999 0 0 1 16 31Z"/>`),
		g.Group(children),
	)
}