package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M832 832V704q0-14-9-23t-23-9H608V480q0-14-9-23t-23-9H448q-14 0-23 9t-9 23v192H224q-14 0-23 9t-9 23v128q0 14 9 23t23 9h192v192q0 14 9 23t23 9h128q14 0 23-9t9-23V864h192q14 0 23-9t9-23zm576 64q0-53-37.5-90.5T1280 768t-90.5 37.5T1152 896t37.5 90.5t90.5 37.5t90.5-37.5T1408 896zm256-256q0-53-37.5-90.5T1536 512t-90.5 37.5T1408 640t37.5 90.5T1536 768t90.5-37.5T1664 640zm256 128q0 212-150 362t-362 150q-192 0-338-128H850q-146 128-338 128q-212 0-362-150T0 768t150-362t362-150h896q212 0 362 150t150 362z"/>`),
		g.Group(children),
	)
}