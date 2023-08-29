package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeIncopy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0v512h512V0H0zm198.982 364.218H157.09v-256h41.89v256zM370.97 214.045c-39.103-15.294-80.498-.433-80.76 54.871c.01 51.564 36.072 70.49 80.76 55.484v37.27c-56.684 18.29-123.103-5.78-123.348-92.754c.219-74.086 58.07-110.679 123.348-93.309v38.438z"/>`),
		g.Group(children),
	)
}