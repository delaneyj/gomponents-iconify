package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tla(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m376.317 21.71l-82.846 81.337H0V490.29l460.078-.449v-284.43L512 153.489L376.317 21.711zM37.632 129.755h230.64l-22.999 22.999l11.316 10.99H37.632v-33.99zM423.66 464.797H37.632V197.733h253.954l89.37 86.799l42.704-42.704v222.969zm-30.348-299.84v69.194H361.75v-69.194h-66.766v-31.562h66.766V69.057h31.562v64.338h66.766v31.562h-66.766zM144.457 286.35v-52.199h173.591v52.199h-26.706v-24.279h-40.06v143.244h27.92v31.562h-88.616v-31.562h27.92V262.07h-48.984v24.279h-25.065z"/>`),
		g.Group(children),
	)
}