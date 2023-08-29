package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityHall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M7.743 21.8h3.485v22.4h5.229V25h5.229v19.2h5.228V25h5.23v19.2h5.229V21.8h3.484c.963 0 1.744-.716 1.744-1.6c0-.534-.288-1.004-.727-1.294l.003-.003l-.026-.015l-.045-.027l-16.635-8.609V7.73c3.072 1.412 5.601-1.02 9.585.442V2.601c-3.986-1.462-6.514.968-9.585-.443V1.8c0-.443-.389-.8-.871-.8s-.87.357-.87.8v8.452L6.795 18.859l-.045.027l-.025.017v.003c-.437.29-.724.761-.724 1.294c-.001.884.78 1.6 1.742 1.6zm1.742 24.001L6 49h36.602l-3.487-3.199z"/>`),
		g.Group(children),
	)
}