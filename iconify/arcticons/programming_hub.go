package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingHub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.153 43.5V23.364m0-12.054V4.5h12.415a12.892 12.892 0 1 1 0 25.784H14.153"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.402 17.485l-10.059 5.807V11.677Zm-10.264-5.808l-11.95-.196m12.154 4.099H10.931m10.241 3.586H8.54m12.495 4.126H10.452"/>`),
		g.Group(children),
	)
}