package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m745 1053l148 148l-149 149zm-1-611l149 149l-148 148zM614 1666l464-464l-306-306l306-306l-464-464v611L359 482l-93 93l320 321l-320 321l93 93l255-255v611zm719-770q0 209-32 365.5t-87.5 257T1073 1681t-181.5 86.5T672 1792t-219.5-24.5T271 1681t-140.5-162.5t-87.5-257T11 896t32-365.5t87.5-257T271 111t181.5-86.5T672 0t219.5 24.5T1073 111t140.5 162.5t87.5 257t32 365.5z"/>`),
		g.Group(children),
	)
}