package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Workit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.148 18.452a2.415 2.415 0 0 1 4.83 0h0m-.004 11.356a2.415 2.415 0 0 1-4.82 0m4.822-11.251v11.207m-4.827-11.206v11.186m-5.634-9.979a1.76 1.76 0 1 1 3.52 0h0m-.002 8.278a1.76 1.76 0 0 1-3.514 0m3.515-8.202v8.17m-3.519-8.169v8.154m31.114-8.224a1.76 1.76 0 1 1 3.52 0h0m-.002 8.278a1.76 1.76 0 0 1-3.515 0m3.516-8.202v8.17M37.63 19.85v8.154m-6.822-9.667a2.415 2.415 0 0 1 4.83 0h0m-.004 11.467a2.415 2.415 0 0 1-4.82 0m4.822-11.364v11.21m-4.827-11.208v11.186m-11.956-6.993h10.079m0 2.52H18.853m24.382-2.523c.7.002 1.265.57 1.265 1.27h0a1.27 1.27 0 0 1-1.265 1.269M4.774 25.17h-.005A1.27 1.27 0 0 1 3.5 23.902h0c0-.7.568-1.269 1.27-1.269h.004"/>`),
		g.Group(children),
	)
}