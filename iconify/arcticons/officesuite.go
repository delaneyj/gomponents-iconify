package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Officesuite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.807 14.858L5.5 24.053l5.277 9.142L16.084 24l-5.277-9.142z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.084 24l-5.307 9.195h10.554L26.639 24H16.084zm-5.277-9.142L16.084 24h10.555l-5.278-9.142H10.807z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.669 5.662l-5.308 9.196L26.639 24l5.307-9.195l-5.277-9.143zm5.277 9.143L26.639 24h10.554l5.307-9.195H31.946z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.669 5.662l5.277 9.143H42.5l-5.277-9.143H26.669zM26.639 24l-5.308 9.195l5.277 9.143l5.308-9.196L26.639 24z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.916 33.142l-5.308 9.196h10.555l5.307-9.196H31.916zM26.639 24l5.277 9.142H42.47L37.193 24H26.639z"/>`),
		g.Group(children),
	)
}