package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42.5 16.3869C41.5521 14.0859 40.1874 12.0006 38.5 10.225C34.8561 6.39055 29.7072 4 24 4C12.9543 4 4 12.9543 4 24H14L19 32L28 14L35 24H44C44 35.0457 35.0457 44 24 44C18.5491 44 13.6075 41.8194 10 38.2829C8.17976 36.4985 6.69917 34.3688 5.66417 32"/>`),
		g.Group(children),
	)
}