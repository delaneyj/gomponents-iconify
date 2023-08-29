package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#61B2E4" d="M36 24.644L29.416 16V4h-5.505v12c0 6.028 4.004 19.958 4.004 19.958L16 67.98h40L44.085 35.96S48.089 22.027 48.089 16V4h-5.506v12L36 24.644z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M36 24.644L29.416 16V4h-5.505v12c0 6.028 4.004 19.958 4.004 19.958L16 67.98h40L44.085 35.96S48.089 22.027 48.089 16V4h-5.506v12L36 24.644zm-4.318 11.314h8.644"/>`),
		g.Group(children),
	)
}