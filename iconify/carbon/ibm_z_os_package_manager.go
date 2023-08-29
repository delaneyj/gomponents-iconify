package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmZOsPackageManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 9v10.172l-2.586-2.586L11 18l5 5l5-5l-1.414-1.414L17 19.172V9h-2zm4.758 17.65L16 28.842L5 22.426V18H3v5c0 .355.189.685.496.864l12 7a.999.999 0 0 0 1.008 0l4.282-2.498l-1.028-1.716zm8.746-18.514l-4.269-2.49l-1.029 1.715L27 9.574v12.852l-3.787 2.209l1.029 1.715l4.262-2.486c.307-.179.496-.509.496-.864V9c0-.355-.189-.685-.496-.864zM5 9.574l11-6.417l3.751 2.188L20.78 3.63l-4.276-2.494a.999.999 0 0 0-1.008 0l-12 7A1.001 1.001 0 0 0 3 9v5h2V9.574z"/>`),
		g.Group(children),
	)
}