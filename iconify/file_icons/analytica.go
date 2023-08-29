package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analytica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M296.879 205.682v-41.214H351v287.614h-54.121v-46.458s-22.663 52.353-99.128 53.466c-68.589 1-151.695-43.916-152.604-149.91c-.886-103.22 77.749-150.577 148.697-151.396c44.598-.514 86.146 20.382 103.035 47.898zm.22 103.06c0-79.002-82.798-128.647-148.679-89.146s-65.88 138.79 0 178.291s148.68-10.143 148.68-89.145zm111.07-204.308V512L512 327.095V0H188.798L0 104.434h408.17z"/>`),
		g.Group(children),
	)
}