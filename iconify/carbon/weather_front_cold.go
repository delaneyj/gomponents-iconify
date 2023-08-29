package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherFrontCold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.166 2a13.029 13.029 0 0 0-12.542 9.58l-1.177 4.315A11.024 11.024 0 0 1 3.835 24H2v2h1.834c.298 0 .592-.017.886-.037l7.929 2.974A1 1 0 0 0 14 28v-6.908a12.931 12.931 0 0 0 .912-1.313l6.363-1.818a1 1 0 0 0 .506-1.586l-3.997-4.995a10.927 10.927 0 0 1 1.152-2.347l5.748 1.916A1 1 0 0 0 26 10V4.225A10.896 10.896 0 0 1 28.166 4H30V2zM8.336 25.183A13.047 13.047 0 0 0 12 23.094v3.463zm7.74-7.816c.11-.31.213-.625.3-.946l.733-2.685l2.171 2.715zM24 8.612l-3.741-1.247A11.071 11.071 0 0 1 24 4.831zM10 4v4.586L3.414 2L2 3.414L8.586 10H4v2h8V4h-2z"/>`),
		g.Group(children),
	)
}