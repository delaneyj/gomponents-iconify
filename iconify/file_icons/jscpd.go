package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jscpd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M68.229 200.902H258.4v19.505H68.23v-19.505zm58.514 48.762h287.695v19.504H126.743v-19.504zm14.629 58.514h224.304v19.505H141.372v-19.505zm-53.639 58.514h224.305v19.505H87.733v-19.505zm-9.752 58.515h312.076v19.504H77.981v-19.504zm47.192-302.324l58.155-102.036l-21.182-12.073l-53.836 94.458l-58.866-93.36L28.82 22.878l63.057 100.006H0v380.343h512V122.883H125.173zm357.57 351.085H29.257V152.14h453.486v321.828z"/>`),
		g.Group(children),
	)
}